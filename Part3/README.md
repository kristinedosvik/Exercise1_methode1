# Reasons for concurrency and parallelism


To complete this exercise you will have to use git. Create one or several commits that adds answers to the following questions and push it to a repository to complete the task. Remember to also
 submit your answers to Blackboard

When answering the questions, remember to use all the resources at your disposal. Asking the internet isn't a form of "cheating", it's a way of learning.

 ### What is concurrency? What is parallelism? What's the difference?
 Concurrency:At flere operasjoner kjører uten at man helt sikkert kan vite hvilke operasjoner som kjører til enhver tid.
Paralellisme: Flere operasjoner kjører samtidig.
Forskjellen: I et parallelt program kjører operasjonene samtidig. I et program som bruker concurrency skjer ikke operasjonene samtidig, men de ulike trådene veksler på å utføre operasjonene
 sine, slik at det virker som om programmet kjører parallelt. 
 
 ### Why have machines become increasingly multicore in the past decade?
Maskinen blir da mer effektiv, og maskinens "speed" øker. I tillegg minker varme og effektforbruket til maskinen.

 
 ### What kinds of problems motivates the need for concurrent execution?
 (Or phrased differently: What problems do concurrency help in solving?)
 > Hvis vi vil kjøre to seperate prosesser samtidig kan concurrency være nyttig.
 
 ### Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
 Filen kan bli vanskeligere å kode mtp. at man må sette visse betingelser for at programmet skal gjøre riktig oppgave til riktig tid.
Det kan bli vanskeligere å lese også fordi man ikke alltid vet når programmet bytter fra tråd til tråd.
 
 ### What are the differences between processes, threads, green threads, and coroutines?
Prosesser: Kjører i separerte minnerom. En prosess kjører et program.
Tråder: Kjører i et felles minnerom dersom de kommer fra samme prosess.
Disse kjører en sekvens i en prosess
Grønne tråder: Er brukerlevel tråd som blir planlagt av brukerens og ikke kernelen.
Coroutiner: datamaskinprogram som sørger for at rutinenr gjennomføres helt før maskinen
begynner å jobbe på en annen tråd. (?)
 
 ### Which one of these do `pthread_create()` (C/POSIX), `threading.Thread()` (Python),
### `go` (Go) create?
Alle funksjonene lager en tråd i programmet i hhv. C, python og go.
 
 
 ### How does pythons Global Interpreter Lock (GIL) influence the way a python Thread behaves?
 > GIL er en lås som kun lar én tråd "holde" pythons interpreteren. Dermed kan kun en tråd
kjøre om gangen
 
 ### With this in mind: What is the workaround for the GIL (Hint: it's another module)?
 > *Man kan bruke en threading module
 
 ### What does `func GOMAXPROCS(n int) int` change? 
 > Forandrer hvor mange CPUer som maksimalt kan operere samtidig.
