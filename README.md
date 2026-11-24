# Design Add and Search Words

> NeetCode 150 · Tries · Blind 75 — a trie where a literal letter follows one child and a '.' tries all 26; the wildcard turns lookup into bounded DFS. Go, stdlib only, tested.

## Problem

A dictionary supporting `AddWord` and `Search`, where search patterns may contain `.` matching exactly one arbitrary letter — `".ad"` hits "bad", "dad", "mad".

## Approach

Plain trie storage (dsa-185); all the novelty is in search. A literal character descends into one child as usual. A `.` branches: recursively try *every* existing child at that depth, succeeding if any continues to a match. The recursion depth is the pattern length, and the fan-out only happens at dots, so all-literal searches stay O(len) while a worst-case all-dots pattern degrades toward visiting the whole trie — the honest price of wildcards. "Exactly one letter, never zero" is the spec detail the `"...."` test pins. This dot-matching is a two-symbol preview of regex matching (coming at dsa-263), minus the `*`.

**Complexity:** AddWord O(len); Search O(len) without dots, up to O(26^dots · len) with.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
