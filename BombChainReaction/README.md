# Bomb Chain Reaction

Given a list of Bombs with a coordinates, `x, y`, and radius, `r`, determine the number of bombs that will detonate if we detonate a given start bomb.

## Clarifications

* Exploding bomb A will trigger bomb B if the euclidean distance is less than bomb A's radius.
* The start bomb is null:
    return 0
* The list is null or empty:
    return 0
* The start bomb may not be in the list
