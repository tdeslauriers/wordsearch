## Word search

My dad made a wordsearch app in the 80's and it seemed like a fun *puzzle* to noodle on.  

1. This reads a dictionary text file.
1. Randomly selects 20 words.
1. Creates a matrix with either 20 positions, or adjusts for a word length larger than 20 letters.
1. Uses recursion (random positions + orientations) to cycle thru trying to place the word in the puzzle.
1. Adds in randomly chose letters to empty spaces.
1. Outputs to console.

###
**Notes:**

1. At first, I was sorting the word list longest to shortest.  
* This makes the brute-forcing go faster, but it also felt too deterministic omitting sorting never failed to produce a matrix, so I removed it.
2. A puzzle will be successfully produced about 75% of the time if you make the puzzle the size of the longest word, instead of default to 20 rows/columns.
2. I used cryptographic random numbers for no reason; I am sure math/rand was fine for this purpose.
1. For a 20 row/column puzzle, 25 appears to be the highest number of words it will reliable resolve for.  At 27, it becomes 50/50.  

## Sample Output

```
[Y S K C C C D P E N I R A G R A M P R X]
[K F O X Z V H U C R Y S T A L L U R G Y]
[U N M I N I S T R A N T U H W W I T P F]
[U N E X P R E S S A B L E M I M R H X G]
[D O Y D E T A N I M O N G O C D J N C A]
[K F M M O R H K X D T W O D L N H W M H]
[X J O S H E T E R O D O N T I D A E S U]
[Y T I C I N O T O N O M O J F K D V Q D]
[F T Z Q C B Q L H S F W N V I P E V N S]
[O N I P S J E V U B O M E Q U O N N Z H]
[D Z W V B T G S W X E O Z K W Q E S D I]
[E T O H N V P Y H T M P P M K P E S S K]
[F I U E F E B N I Q J E G M J P O W J A]
[C R U Y C R X M K Z Y G Q N L U K I D R]
[M Q Q T B R A L L D J S J Z F K A R I E]
[P F F T I N D O E U R O P E A N A V L E]
[Y U Y S Y J G N D T S N I L B O G K Q S]
[L N Q D V L C I L Y T C A D O R E T P F]
[O G O V J Y T I L A E R A I I Z K P Y X]
[S D R A I N A L H I S O S M O T I C V U]
```

1. AREALITY
2. COGNOMINATED
3. CRYSTALLURGY
4. DYNAMITE
5. GOBLINS
6. HETERODONTIDAE
7. HOTE
8. INDOEUROPEAN
9. IOP
10. ISOSMOTIC
11. LANIARDS
12. MARGARINE
13. MEQUON
14. MONOTONICITY
15. PTERODACTYLIC
16. SHIKAREES
17. SUSPECTFUL
18. UNEXPRESSABLE
19. UNMINISTRANT
20. WICLIF