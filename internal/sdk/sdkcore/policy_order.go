package forticlient

import (
	"fmt"
	"net/http"
	"slices"
	"strings"
)

// EnsurePolicyOrder skips moving only when the order is confirmed as satisfied.
// A failed order check still allows the move. Empty anchors leave order unmanaged.
// The caller must hold the resource's write lock throughout this method.
func (c *FortiSDKClient) EnsurePolicyOrder(collectionURL, key, before, after string) error {
	if before == "" && after == "" {
		return nil
	}
	anchor, _, err := getPolicyOrderTarget(key, before, after)
	if err != nil {
		return err
	}
	inSync, err := c.checkPolicyOrder(collectionURL, key, anchor, before != "")
	if err == nil && inSync {
		return nil
	}
	if err := c.MovePolicy(collectionURL, key, before, after); err != nil {
		return fmt.Errorf("move policy %q: %w", key, err)
	}
	return nil
}

func (c *FortiSDKClient) checkPolicyOrder(collectionURL, key, anchor string, before bool) (bool, error) {
	keys, err := c.ListPolicyOrder(collectionURL)
	if err != nil {
		return false, fmt.Errorf("read policy order: %w", err)
	}
	policyIndex, anchorIndex := slices.Index(keys, key), slices.Index(keys, anchor)
	if policyIndex < 0 {
		return false, fmt.Errorf("policy %q was not found in collection %s", key, collectionURL)
	}
	if anchorIndex < 0 {
		return false, fmt.Errorf("ordering anchor %q was not found in collection %s", anchor, collectionURL)
	}
	if before {
		return policyIndex < anchorIndex, nil
	}
	return policyIndex > anchorIndex, nil
}

// ListPolicyOrder returns primary keys in the order supplied by the collection API.
// It must not use read(), which returns only the first entry of an array.
func (c *FortiSDKClient) ListPolicyOrder(collectionURL string) ([]string, error) {
	result, err := sendRequests(c, &InputModel{
		HTTPMethod: http.MethodGet,
		URL:        collectionURL,
	})
	if err != nil {
		return nil, err
	}
	entries, ok := result["data"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("policy collection %s must return a data array", collectionURL)
	}
	keys := make([]string, 0, len(entries))
	seen := make(map[string]bool, len(entries))
	for index, entry := range entries {
		object, ok := entry.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("policy collection entry %d must be an object", index)
		}
		key, ok := object["primaryKey"].(string)
		if !ok || strings.TrimSpace(key) == "" {
			return nil, fmt.Errorf("policy collection entry %d has no non-empty string primaryKey", index)
		}
		if seen[key] {
			return nil, fmt.Errorf("policy collection contains duplicate primaryKey %q", key)
		}
		seen[key] = true
		keys = append(keys, key)
	}
	return keys, nil
}

// MovePolicy moves an existing entry before or after exactly one other entry.
// Callers are responsible for holding the resource's write lock.
func (c *FortiSDKClient) MovePolicy(collectionURL, key, before, after string) error {
	anchor, parameter, err := getPolicyOrderTarget(key, before, after)
	if err != nil {
		return err
	}
	_, err = sendRequests(c, &InputModel{
		HTTPMethod: http.MethodPost,
		URL:        collectionURL,
		HeadParams: map[string]interface{}{
			"action":  "move",
			"key":     key,
			parameter: anchor,
		},
	})
	return err
}

func getPolicyOrderTarget(key, before, after string) (string, string, error) {
	if strings.TrimSpace(key) == "" {
		return "", "", fmt.Errorf("policy key must not be empty")
	}
	if (before == "") == (after == "") {
		return "", "", fmt.Errorf("exactly one of before_policy or after_policy must be provided")
	}
	anchor := before
	parameter := "before"
	if after != "" {
		anchor, parameter = after, "after"
	}
	if strings.TrimSpace(anchor) == "" || anchor == key {
		return "", "", fmt.Errorf("ordering anchor must be non-empty and different from policy %q", key)
	}
	return anchor, parameter, nil
}
