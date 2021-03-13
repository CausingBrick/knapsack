# Knapsack Problem

## 01 Knapsack Problem

### Introduction

In the 01 Knapsack problem, we are given a knapsack of fixed capacity C. We are also given a list of N objects, each having a weight W(I) and profit P(I). We can put any subset of the objects into the knapsack, as long as the total weight of our selection does not exceed C. We desire to maximize our total profit, which is the sum of the profits of each object we put into the knapsack.

Thus, a solution of the 01 Knapsack problem is a subset S of the N objects for which the weight sum is less than or equal to C, and which maximizes the total profit.

### Greedy Strategies

#### Design

We can easily come up with three different greedy strategies which are for selecting suitable items to put in the knapsack, although maybe they are not enough fast.

1. Pick the lightest item in each selection.
2. Pick the best profit item in each selection.
3. Pick the item has highest ratio of profit and weight in each selection.

#### Benchmark

