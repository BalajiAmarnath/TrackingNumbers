# TrackingNumbers

Given, a table of transaction records with a starting tracking number, ending tracking number, status code and transfer code, 
the table must always accurately represent the most recent data (status code and transfer code) for that tracking number. 

<h4>Input</h4>
<p>Each input case begins with a single line that is a character string naming the test case. This string contains at most 80 
characters. The name "END" marks the end of the input. Following this will be 1 to 100 lines of the form "A B S T", where A, B,
and T are integers in the range 1 to 231-1, S is an uppercase letter, and A<=B. These lines are, in the order they are to be 
applied, the tracking number transactions to be recorded, where A is the start of the tracking number range, B is the end of
the tracking number range, S is the status code, and T is the transfer code. The list of tracking number transactions is 
terminated by a line containing only a 0 (zero) character.</p>

<h4>Output</h4>
<p>For each input case, echo the test case name to the output on a line by itself, followed by the resulting minimal-rows 
tracking number table that results after all tracking number transactions have been applied.</p>
