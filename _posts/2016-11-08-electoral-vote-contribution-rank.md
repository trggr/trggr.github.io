---
layout: post
---

#### Electoral Votes

Interesting how the states' populations translated into
electoral votes. I got the [data](http://state.1keydata.com/state-electoral-votes.php)
and built a spreadsheet to see which states' populations have more
weight in the outcome of the voting.

For example, Wyoming with its population of 586K people has three
electoral votes. It means every 195K people of Wyoming as much voting clout
as 722K people of Texas. It's 3.7 times more. 
It's also 2.6 times more than average.

	 State         |  Pop   |EV|Ratio | Rank
	--------------------------------------- 
	 Wyoming       |  586107| 3|195369| 1 
	 Vermont       |  626042| 3|208681| 2 
	 Alaska        |  738432| 3|246144| 3 
	 North Dakota  |  756927| 3|252309| 4 
	 Rhode Island  | 1056298| 4|264075| 5 
	 South Dakota  |  858469| 3|286156| 6 
	 Delaware      |  945934| 3|315311| 7 
	 Maine         | 1329328| 4|332332| 8 
	 New Hampshire | 1330608| 4|332652| 9 
	 Montana       | 1032949| 3|344316|10 
	 Hawaii        | 1431603| 4|357901|11 
	 West Virginia | 1844128| 5|368826|12 
	 Nebraska      | 1896190| 5|379238|13 
	 Idaho         | 1654930| 4|413733|14 
	 New Mexico    | 2085109| 5|417022|15 
	 Nevada        | 2890845| 6|481808|16 
	 Kansas        | 2911641| 6|485274|17 
	 Arkansas      | 2978204| 6|496367|18 
	 Mississippi   | 2992333| 6|498722|19 
	 Utah          | 2995919| 6|499320|20 
	 Connecticut   | 3590886| 7|512984|21 
	 Iowa          | 3123899| 6|520650|22 
	 Alabama       | 4858979| 9|539887|23 
	 South Carolina| 4896146| 9|544016|24 
	 Minnesota     | 5489594|10|548959|25 
	 Kentucky      | 4425092| 8|553137|26 
	 Oklahoma      | 3911338| 7|558763|27 
	 Oregon        | 4028977| 7|575568|28 
	 Wisconsin     | 5771337|10|577134|29 
	 Louisiana     | 4670724| 8|583841|30 
	 Washington    | 7170351|12|597529|31 
	 Tennessee     | 6600299|11|600027|32 
	 Maryland      | 6006401|10|600640|33 
	 Indiana       | 6619680|11|601789|34 
	 Colorado      | 5456574| 9|606286|35 
	 Missouri      | 6083672|10|608367|36 
	 Massachusetts | 6794422|11|617675|37 
	 Michigan      | 9922576|16|620161|38 
	 Arizona       | 6828065|11|620733|39 
	 Georgia       |10214860|16|638429|40 
	 New Jersey    | 8958013|14|639858|41 
	 Pennsylvania  |12802503|20|640125|42 
	 Illinois      |12859995|20|643000|43 
	 Virginia      | 8382993|13|644846|44 
	 Ohio          |11613423|18|645190|45 
	 North Carolina|10042802|15|669520|46 
	 New York      |19795791|29|682613|47 
	 Florida       |20271272|29|699009|48 
	 California    |39144818|55|711724|49 
	 Texas         |27469114|38|722871|50 

Few interesting things to make the calculations in Excel and output the
results in fixed length format, so I can copy and paste the table
into markdown doc.

I've made an Excel table Population with columns:

	State(PK)		  - state name
	Pop               - state population
	Votes			  - electoral votes
	PopVoteRatio	  =[Pop]/[Votes]
	PopVoteRatioRank  =RANK([@PopVoteRatio],[PopVoteRatio],1)
	StateLen		  =LEN([State])
	StatePadded		  =[State]&REPT(" ", MAX([StateLen])-LEN([State]))
	Export            =[StatePadded]                 &"|"&
                       TEXT([Pop],"????????")        &"|"&
                       TEXT([Votes],"??")            &"|"&
                       TEXT([PopVoteRatio],"??????") &"|"&
                       TEXT([PopVoteRatioRank], "??")



