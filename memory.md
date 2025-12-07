# memory 

- An application creates a process.
- The OS creates the process ..
- A process the unit of execution

- A process contains some memory 

- Text/Code Segment --> The binary is loaded into the text segment
- Data Segment -> RO,BSS --> all constants, global variables are stored in Data Segment
- Stack Memory -> 2mb on linux machines, 1 mb on windows machine, or 4mb or 8mb (depends on OS and settings)
- Heap Memory --> any amount of memory (subject to availability)

- If you create a variable and assign data, to store the on stack or heap is decided by the compiler, the concept is called as escape analysis that is done by the compiler
