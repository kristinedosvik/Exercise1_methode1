Part1: Thinking about elevators
---------------------------

Not for handing in, just for thinking about. Talk to other groups, assistants, or even people who have taken the course in previous years.

Brainstorm some techniques you could use to prevent a user from being hopelessly stranded, waiting for an elevator that will never arrive. Think about the [worst-case](http://xkcd.com/748/)
 behaviour of the system.
 - What if the software controlling one of the elevators suddenly crashes?
 Da bør sw som kontrollerer andre heiser ta over
 - What if it doesn't crash, but hangs?
Hvis den henger bør man restarte og andre heiser bør ta over
 - What if a message between machines is lost?
Man bør ha en send og noe som sjekker om beskjeden ble levert, ble den ikke levert må man sende på nytt. Evt. restarte.
 - What if the network cable is suddenly disconnected? Then re-connected?
Hvis nettverket forsvinner må man ha lagret informasjonen på heisene, og deretter sørge for at man sjekker om informasjonen som skulle sendes ble sendt.
 - What if a user of the system is being a troll?
Da må systemet reagere på det på en slik måte den ville reagert på andre brukere.
 - What if the elevator car never arrives at its destination?
Man må sørge for at en heis når destinasjonen, uansett.
