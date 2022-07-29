<template>
  <v-card>
    <v-card-title class="justify-center">League Rules</v-card-title>
  <div>
    <blockquote class="blockquote">
      <p>Below are the proposed rules for a best ball, contract, dynasty fantasy football league. <br>
      The intent is to create a league that is low/no maintenance during the season, but still allows for lots of fun in the off season.</p>
      <p>*Changes in July 2022</p>
      <ul>
            <li>Added clarification on off-season execution: no drops after restructuring window through rookie signing</li>
            <li>Modified rule to allow for being over the cap in future years but requiring all owners to have all years under the cap after the restructuring window every season</li>
            <li>Clarified that owners cannot “eat” cap during a trade.</li>
            <li>Updated Rookie contracts to allow for signing to shorter deals and clarified slope.</li>
            <li>Added commissioner discretion rule</li>
            <li>Updated roster size to 40 based on vote</li>
            <li>Confirmed Matching Tags cannot be traded based on vote</li>
      </ul>
    </blockquote>

    <v-expansion-panels
      v-model="panel"
      multiple
      focusable
    >
      <v-expansion-panel>
        <v-expansion-panel-header class="grey--text">Core League Elements</v-expansion-panel-header>
        <v-expansion-panel-content>
          <br/>
          <h3><u>Best Ball</u></h3>
          <p>Your best performing players are automatically "started". You make no start-sit decisions and there is no weekly waiver wire. The roster is a bit bigger as a result to account for injuries and bye weeks. Most standard best ball leagues don't allow trading (they are just draft exercises), but this league will have trading. Matchups are still head-to-head each week.</p>
          <h3><u>Contract</u></h3>
          <p>Dynasty version of an auction draft. The start-up draft is an auction of all players (including rookies for the first season only). You can offer a player a 1-year, 2-year, 3-year, or 4-year deal. There is a discount for longer contracts (e.g. a $10 million 1-year deal may be worth less than a $19.5 million 2-year deal). There will be a predetermined "slope" for this that is provided below. </p>
          <h3><u>Salary Cap</u></h3>
          <p>There will be a hard salary cap ($200 million/yr) and a minimum contract (1-year, $0.5 million). These loosely match the NFL cap and minimum.Salary cap, dead money, and “cash” cannot be traded or adjusted during a trade (no eating salary to enable trades).</p>
          <h3><u>Yearly Draft and Auction</u></h3>
          <p>In the offseason, all players not on rosters plus those that have finished their contracts will go into a free agent pool. Owners will bid in a similar way as the start-up draft. This will occur in the spring.</p>
          <p>There will be a rookie draft every year (except 2021) that follows a 1-12 order for every round based on last season's result (Last place selects first, and so on, non-snake draft). Rookie contract values are predetermined by pick and follow the same slope as other contracts and can be 1-4 years in length.</p>
          <p>There will be a low-intensity mid-season free agent auction (Dutch Auction) to place any free agents onto rosters for the remainder of the season.</p>
        </v-expansion-panel-content>
      </v-expansion-panel>

      <v-expansion-panel>
        <v-expansion-panel-header class="grey--text">Roster Composition</v-expansion-panel-header>
        <v-expansion-panel-content>
          <br/>
          <p>We want to introduce some new elements with the starting line-up to make it more interesting.
             Remember, you don't select a line-up each week, but the slots are filled with your best players to maximize scoring. 
             It's also heavier on flex spots to make it easier to have a valid line-up each week. The effect of the roster is to increase value for QBs and TEs (rather than doing that through scoring only).</p>
          <ul>
            <li>1 QB</li>
            <li>1 RB</li>
            <li>2 WR</li>
            <li>2 TE</li>
            <li>1 Superflex (QB/RB/WR/TE)</li>
            <li>2 Flex (RB/WR)</li>
          </ul>
          <p></p>
          <p>No kickers or defenses. Obviously. They are stupid in all leagues, but especially stupid in this format.</p>
          <p>Maximum roster size is 35 players based on starting as a 12-team league.
             However, you may not want to roster that many due to the salary cap and minimum salary.
              The maximum roster may adjust down or up based on teams in the league or how we find things work in year 1.</p>
        </v-expansion-panel-content>
      </v-expansion-panel>

      <v-expansion-panel>
        <v-expansion-panel-header class="grey--text">Scoring</v-expansion-panel-header>
        <v-expansion-panel-content>
          <br/>
          <p>Same as roster construction, we want some new elements to keep things interesting! Goal is to maintain balance but create much higher scores, 
            so I'm sure we'll tweak things a bit once we can dig into projections a bit more (still on to-do list). We'll use decimal scoring and scoring is additive 
            (e.g. 0.5 PPR and 0.2 points per target = 0.7 total PPR).</p>
          <v-simple-table>
            <template v-slot:default>
              <thead>
                 <tr>
                    <th>Passing</th>
                    <th>Rushing</th>
                    <th>Receiving</th>
                    <th>Miscellaneous</th>
                 </tr>
              </thead>
              <tbody>
                <tr>
                    <td>0.04 pts per passing yard (1 pt per 25 yards)</td>
                    <td>0.1 pts per rushing yard</td>
                    <td>0.1 pts per receiving yard</td>
                    <td>0.04 pts per kickoff return yard</td>
                 </tr>
                 <tr>
                    <td>0.1 pts per complete pass</td>
                    <td>0.1 pts per rushing attempt</td>
                    <td>0.5 pts per reception</td>
                    <td>0.5 pts per reception</td>
                 </tr>
                 <tr>
                    <td>-0.1 pts per incomplete pass</td>
                    <td>6 pts per rushing TD</td>
                    <td>0.2 pts per receiving target</td>
                    <td>6 points per kickoff return TD</td>
                 </tr>
                 <tr>
                    <td>6 pts per passing TD</td>
                    <td>2 pts per rushing 2 point conversion</td>
                    <td>6 pts per receiving TD</td>
                    <td>6 points per punt return TD</td>
                 </tr>
                 <tr>
                    <td>-5 pts per INT</td>
                    <td>3 points for 200+ yard Rushing Game</td>
                    <td>2 pts per receiving 2 point conversion</td>
                    <td>6 points per fumble recovered for a TD</td>
                 </tr>
                 <tr>
                    <td>2 pts per passing 2 point conversion</td>
                    <td></td>
                    <td>3 points for 200+ yard Receiving Game</td>
                    <td>-1 points per fumble</td>
                 </tr>
                 <tr>
                    <td>-0.5 pts per SACK taken</td>
                    <td></td>
                    <td></td>
                    <td>-2 points per fumble lost</td>
                 </tr>
                 <tr>
                    <td>3 points for 400+ yard Passing Game</td>
                    <td></td>
                    <td></td>
                    <td></td>
                 </tr>
              </tbody>
            </template>
          </v-simple-table>
        </v-expansion-panel-content>
      </v-expansion-panel>

      <v-expansion-panel>
        <v-expansion-panel-header class="grey--text">Season Milestones / Deadline Summary</v-expansion-panel-header>
        <v-expansion-panel-content>
          <br/>
          <h3><u>Roughly Mid-May</u></h3>
          <p>The league year begins with the submission of all restructures and drops 2 weeks before the offseason Free Agent Auction. At this point, all teams must be under the salary cap in all league years (excluding rookie picks as these players are not signed yet).</p>
          <h4>Important:</h4>
          <ul>
            <li>No restructuring is allowed until the next offseason</li>
            <li>No drops are allowed until after the rookie draft.</li>
            <li>No trades are allowed until after the free agent auction.</li>
          </ul>
          <p></p>
          <h3><u>Roughly Late May</u></h3>
          <p>Free Agent Auction</p>
          <h3><u>Early June</u></h3>
          <p>All contract lengths for free agents signed during the auction must be confirmed.<b>Trades are allowed after contract lengths are confirmed</b></p>
          <h3><u>Mid-July</u></h3>
          <p>Rookie Draft (slow draft by email). <b>Trades are allowed, but drops are still not allowed</b></p>
          <h3><u>Early August</u></h3>
          <p>Rookie contracts must be confirmed within 1 week of the end of the draft.<b> Drops are now allowed.</b></p>
          <h3><u>NFL Week 1</u></h3>
          <p>Contracts start to get "spent" against. 1/10 the yearly value of the contract per week through 10.</p>
          <h3><u>NFL Week 6</u></h3>
          <p>Mid-Season Free Agent auction (dutch auction)</p>
          <h3><u>NFL Week 10</u></h3>
          <p>Transaction deadline for that season. No transactions (drops, trades) allowed until the end of the fantasy playoffs</p>
          <h3><u>NFL Week 18</u></h3>
          <p>All transactions now allowed</p>
        </v-expansion-panel-content>
      </v-expansion-panel>

      <v-expansion-panel>
        <v-expansion-panel-header class="grey--text">Contracts and Free Agent Auction</v-expansion-panel-header>
        <v-expansion-panel-content> <br/>
          <p>The minimum contract for a player is a 1 year, $0.5 million. 
            This is the opening bid in the auction when a player is nominated. 
            Nominations go in reverse order from the draft order and owners can pass on a nomination (still able to nominate in future rounds). 
            Once all owners pass, the draft is over.</p>
            <p></p>
            <p>Once a player is nominated, you'll be given 4 bidding options:</p>
            <ul>
            <li>1 Year, $1 Million</li>
            <li>2 Year, $1.8 Million ($0.90 Million per Year)</li>
            <li>3 Year, $2.6 Million ($0.85 Million per Year)</li>
            <li>4 Year, $3.2 Million ($0.80 Million per Year)</li>
          </ul>
          <p></p>
          <p>Once any bid of those 4 is made, the next increment is a $1.5 Million, 1 Year deal and the associated multi-year deals to that.
             This continues until there is no more bidding. 
             There will be a way to skip ahead so we aren't waiting an hour for a high value player to be bid up.</p>
          <p>This slope for contract lengths will hold for all salaries (e.g. $10 Million, 1 Year deal will be equivalent to a $32 Million, 4 Year deal).</p>
          <p>We will plan on making a one-time adjustment to these slopes based on the first few years of the league. 
            I am sure we don’t have the slopes dialed in correctly. My goal is that there is a dispersion of contracts across the lengths 
            (vs. all 1 year or all 4 year). We’ll work on the analysis after year 2 and make the adjustment for contracts signed in the year 3 offseason free agent draft 
            (obviously not retroactive to contracts signed already). Slope options will be voted with a simple majority needed to pass.</p>
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header class="grey--text">Transaction Deadline</v-expansion-panel-header>
        <v-expansion-panel-content> <br/>
          <p>The league will have a transaction deadline at kickoff of the Sunday, 1pm games of Week 10 of the NFL season. 
            After that point, there will be no allowed trades or roster cuts until after the end of the season.</p>
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header class="grey--text">Restricted Free Agent Matching Tag</v-expansion-panel-header>
        <v-expansion-panel-content> <br/>
          <p>We are also introducing a restricted free agent mechanism (matching tag) to allow owners to more easily retain players. 
            If the player up for auction was on your team, you have the option to match the highest bid made after the auction is complete. 
            You have the choice to change the contract length within the same value. 
            Before you make your choice, the winning bidder can increase their bid, but they are locked into that amount even if you decline to match.</p>
          <p></p>
          <ul>
            <li><b>Example 1</b> - You have a player that you had signed on your team with a 4-year deal. In the offseason, their contract expires, and they go into the free agent auction pool. Another owner wins the bidding with an offer of 1 year, $15 Million. You have the choice to match that contract (following the same slope as normal). </li>
            <li><b>Example 2</b> - You have a rookie you drafted. In the offseason, their contract expires, and they go into the free agent auction pool. Another owner wins the bidding with an offer of 4 years, $32 Million. Before you decide if you want to match, they decide to increase that bid to 3 years, $30 Million. You now choose to match or not. </li>
          </ul>
          <p></p>
          <p>Per league vote, matching tags cannot be traded. They are locked based on ownership at the transaction deadline the season prior.</p>
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header class="grey--text">Salary Cap and Contract Guarantees</v-expansion-panel-header>
        <v-expansion-panel-content> <br/>
          <p>The salary cap is a hard cap at $200 Million (no luxury tax nonsense). Because of the slope, you should always have the highest cap hit in the upcoming season (year 2 should usually be less than year 1). However, your rookie picks will need to be signed (these values are pre-determined).</p>
          <p>Contracts are 50% guaranteed. If you cut a player in the offseason, you eat 50% of the money due the next season in the next season and 50% of the balance in the year after. Money already spent on the player is obviously 100% spent with the maximum being hit at the trade deadline which is week 10. Some examples:</p>
          <ul>
            <li><b>Example 1</b> - You have a player signed to a 4-year, $40 Million Deal signed in the 2020 offseason. After the 2021 season, you cut them. You will have spent $10 million in 2021, you will be assessed $5 Million against the 2022 cap (50% of $10 Million) and $10 million against the 2023 cap (50% of $10 million for 2023 and $10 million for 2024).</li>
            <li><b>Example 2</b> - You have a player signed to a 2 year, $20 Million deal signed in the 2021 season. You cut them in after week 4 of the 2021 season. You will be assessed $7 million against the 2021 cap ($10 Million * 4/10 + $10 Million * 50% * 6/10) and $5 million against the 2022 cap ($10 million * 50%)</li>
          </ul>
          <p></p>
          <p>You can always trade a player to someone with cap space to have them cut the player and eat the cap hit (Brock Osweiler style). However, you cannot trade cap space (all owners have the same hard cap).</p>
          <p>You can restructure the terms of a contract during the offseason with some limitations. This can only be done after the first full season of a multi-year contract and you may not adjust the yearly payout more than 50%. If you make any adjustment, the contract becomes fully guaranteed.</p>
          <ul>
            <li><b>Example 3</b> - You have a player signed to a 4-year, $40 million contract signed in 2021. In the offseason between 2021 and 2022, you restructure the deal to pay $7.5 million in 2022, $7.5 million in 2023, and $15 million in 2024. The contract is now <u>fully guaranteed</u> with the same cap hit rules as above.
              <ul>
                <li>Note: You cannot do 5/5/20 as this would adjust 2024 by 100% which is not allowed (50% max). 5/10/15 would be allowed. </li>
              </ul>
            </li>
          </ul>
          <p></p>
          <p>The commissioner may choose to not allow an owner to drop a player with a fully guaranteed contract if he determines that the drop is mainly intended to make the team worse to improve draft pick position. This falls under the Tanking rule noted below.</p>
          <p>You can never be over the cap within a league year (defined above as starting with the transaction confirmations ahead of the free agent auction and ending at Week 18 of the NFL season). Any transaction that would put you over the cap within the current league year is not allowed. This includes signing free agents, signing rookies, or trades. </p>
          <p>The league has decided that we will allow owners to have signed contracts that go over the cap in future years. However, all owners must be below the cap in all years at the start of the league year. If an owner is not under the cap in any year, they must execute transactions to get under the cap. If completely unable to do this or execute a trade, they forfeit the league year.</p>
          <p>The contract becomes void if a player retires (files retirement paperwork with the NFL). Being cut or injured does not void a contract unless the player then retires.</p>
        </v-expansion-panel-content>
      </v-expansion-panel>
       <v-expansion-panel>
        <v-expansion-panel-header class="grey--text">Rookie Contracts</v-expansion-panel-header>
        <v-expansion-panel-content> <br/>
          <p>The league will have a 5-round rookie draft every season. The order of picks is the reverse of the league finish (league champion has last pick). Only rookies may be taken during the rookie draft.</p>
          <p>As with the NFL, rookie contracts will all be 4-year deals. Values are predetermined based on the draft pick. After the rookie draft, they follow the same rules as the free agent contracts. Picks are forfeited if the team does not have the cap space for the pick or if the owner chooses not to draft a player.</p>
          <p>Rookie contract values have been confirmed per table below. Note that they change from 2022 to 2023 but will stay at 2023 values unless we vote to change. Values are in $ Million/yr for a 1-year deal. Owners can choose to sign picks for multi-year deals, and contracts for rookies follow the same slope as for free agents.</p>
          <v-simple-table>
            <template v-slot:default>
              <thead>
                 <tr>
                    <th>Pick</th>
                    <th>2022</th>
                    <th>2023+</th>
                    <th>Pick</th>
                    <th>2022</th>
                    <th>2023+</th>
                 </tr>
              </thead>
              <tbody>
                <tr>
                    <td>1</td>
                    <td>$24,000,000/yr</td>
                    <td>$12,000,000/yr</td>
                    <td>19</td>
                    <td>$1,750,000/yr</td>
                    <td>$1,500,000/yr</td>
                 </tr>
                 <tr>
                    <td>2</td>
                    <td>$20,500,000/yr</td>
                    <td>$10,500,000/yr</td>
                    <td>20</td>
                    <td>$1,750,000/yr</td>
                    <td>$1,500,000/yr</td>
                 </tr>
                 <tr>
                    <td>3</td>
                    <td>$17,500,000/yr</td>
                    <td>$9,000,000/yr</td>
                    <td>21</td>
                    <td>$1,500,000/yr</td>
                    <td>$1,250,000/yr</td>
                 </tr>
                 <tr>
                    <td>4</td>
                    <td>$14,750,000/yr</td>
                    <td>$7,500,000/yr</td>
                    <td>22</td>
                    <td>$1,500,000/yr</td>
                    <td>$1,250,000/yr</td>
                 </tr>
                 <tr>
                    <td>5</td>
                    <td>$12,250,000/yr</td>
                    <td>$6,250,000/yr</td>
                    <td>23</td>
                    <td>$1,250,000/yr</td>
                    <td>$1,250,000/yr</td>
                 </tr>
                 <tr>
                    <td>6</td>
                    <td>$10,000,000/yr</td>
                    <td>$5,250,000/yr</td>
                    <td>24</td>
                    <td>$1,250,000/yr</td>
                    <td>$1,250,000/yr</td>
                 </tr>
                 <tr>
                    <td>7</td>
                    <td>$8,250,000/yr</td>
                    <td>$4,250,000/yr</td>
                    <td>25</td>
                    <td>$1,000,000/yr</td>
                    <td>$1,000,000/yr</td>
                 </tr>
                 <tr>
                    <td>8</td>
                    <td>$6,750,000/yr</td>
                    <td>$3,500,000/yr</td>
                    <td>26</td>
                    <td>$1,000,000/yr</td>
                    <td>$1,000,000/yr</td>
                 </tr>
                 <tr>
                    <td>9</td>
                    <td>$5,750,000/yr</td>
                    <td>$3,000,000/yr</td>
                    <td>27</td>
                    <td>$1,000,000/yr</td>
                    <td>$1,000,000/yr</td>
                 </tr>
                 <tr>
                    <td>10</td>
                    <td>$4,750,000/yr</td>
                    <td>$2,500,000/yr</td>
                    <td>28</td>
                    <td>$1,000,000/yr</td>
                    <td>$1,000,000/yr</td>
                 </tr>
                 <tr>
                    <td>11</td>
                    <td>$4,000,000/yr</td>
                    <td>$2,250,000/yr</td>
                    <td>29</td>
                    <td>$1,000,000/yr</td>
                    <td>$1,000,000/yr</td>
                 </tr>
                 <tr>
                    <td>12</td>
                    <td>$3,500,000/yr</td>
                    <td>$2,000,000/yr</td>
                    <td>30</td>
                    <td>$1,000,000/yr</td>
                    <td>$1,000,000/yr</td>
                 </tr>
                 <tr>
                    <td>13</td>
                    <td>$3,000,000/yr</td>
                    <td>$1,750,000/yr</td>
                    <td>31</td>
                    <td>$750,000/yr</td>
                    <td>$750,000/yr</td>
                 </tr>
                 <tr>
                    <td>14</td>
                    <td>$2,750,000/yr</td>
                    <td>$1,750,000/yr</td>
                    <td>32</td>
                    <td>$750,000/yr</td>
                    <td>$750,000/yr</td>
                 </tr>
                 <tr>
                    <td>15</td>
                    <td>$2,500,000/yr</td>
                    <td>$1,750,000/yr</td>
                    <td>33</td>
                    <td>$750,000/yr</td>
                    <td>$750,000/yr</td>
                 </tr>
                 <tr>
                    <td>16</td>
                    <td>$2,250,000/yr</td>
                    <td>$1,750,000/yr</td>
                    <td>34</td>
                    <td>$750,000/yr</td>
                    <td>$750,000/yr</td>
                 </tr>
                 <tr>
                    <td>17</td>
                    <td>$2,000,000/yr</td>
                    <td>$1,500,000/yr</td>
                    <td>35</td>
                    <td>$750,000/yr</td>
                    <td>$750,000/yr</td>
                 </tr>
                 <tr>
                    <td>18</td>
                    <td>$2,000,000/yr</td>
                    <td>$1,500,000/yr</td>
                    <td>36</td>
                    <td>$750,000/yr</td>
                    <td>$750,000/yr</td>
                 </tr>
                 <tr>
                    <td></td>
                    <td></td>
                    <td></td>
                    <td>37+</td>
                    <td>$600,000/yr</td>
                    <td>$600,000/yr</td>
                 </tr>
              </tbody>
            </template>
          </v-simple-table>
        </v-expansion-panel-content>
      </v-expansion-panel>
       <v-expansion-panel>
        <v-expansion-panel-header class="grey--text">Mid-Season Free Agent Auction</v-expansion-panel-header>
        <v-expansion-panel-content> <br/>
          <p>Every year there are some surprise players that burst onto the scene. The undrafted free agents, the resurgent veterans, etc. Around week 6, we will have a Dutch auction to allow these players to be acquired (high bidder wins and pays 1 increment higher than the second highest bidder) with bids being the amount you are willing to pay to acquire them for the remainder of the season. The player goes into the free agent auction at the end of the year; however, you will have the ability to use matching tags on these players which will make it easier to retain them.</p>
          <p></p>
          <ul>
            <li><b>Example</b> - You are currently using $190 Million out of your $200 Million salary cap. During the mid-season auction, you bid $10 Million for a player and win the bid. You are now using your full ($200 Million) Salary Cap for the season and gain a free matching tag for that player when they are auctioned the next offseason </li>
          </ul>
        </v-expansion-panel-content>
      </v-expansion-panel>
       <v-expansion-panel>
        <v-expansion-panel-header class="grey--text">Regular Season and Playoff Structure</v-expansion-panel-header>
        <v-expansion-panel-content> <br/>
          <p>The regular season for the league will consist of 14 head-to-head games. Each team will play 3 teams twice per season. The first year this will be randomized. In subsequent years, we will group the teams based on the previous year finish. Teams that finished 1-4 will play each other twice, 5-8, 9-12. This is a simple solution to how to decide on repeat games that drives parity.</p>
          <p>Playoffs will be played from weeks 15-17 of the NFL regular season. The top two regular season finishers based on ESPN seeding (record as first tiebreaker, then points scored) will get a bye for week 15. The 3 plays the 6, the 4 plays the 5. Playoff is single elimination. In week 16, the 1 seed plays the lowest remaining seed and the 2 seed plays the higher (re-seeded).</p>
          <p>There will be a 3-point home-field advantage for the higher seeded team for all rounds.</p>
        </v-expansion-panel-content>
      </v-expansion-panel>
    <v-expansion-panel>
      <v-expansion-panel-header class="grey--text">Buy-In and Payouts</v-expansion-panel-header>
        <v-expansion-panel-content> <br/>
          <p>Buy-in continues to be a source of division within interested people. Some people want a high amount and others have this as a barrier for entry.</p>
          <p>We want this league to have as many teams as possible. We will therefore create a dual buy-in tier system. This will allow for a lower buy-in for more casual owners and a higher buy-in for those that want to gamble.</p>
          <p></p>
          <ul>
            <li><b>Tier 1</b> - $100 yearly buy-in which increases $25 every 2 years. 
              <ul>
                <li>All owners will have to buy-in for at least this amount. This is paid out from playoff results only (exact amounts TBD based on number of owners).</li>
              </ul>
            </li>
            <li><b>Tier 2</b> - Tier 1 plus additional $100 yearly also increasing $25 every 2 years 
              <ul>
                <li>Optional for the bold and strong of heart. This will be paid out from regular season results only (likely 70% for highest finisher of Tier 2 participants and 30% for second highest)</li>
              </ul>
            </li>
          </ul>
          <p></p>
          <p>This dual tier system is to allow for more owners to play, not to be an expected value optimizer. Owners need to decide at the beginning of the league which Tier they want to be in. You’ll be allowed to move up from Tier 1 to Tier 2 with a 2-year delay (if you tell me before the 2021 season you buy in for the higher amount in 2023). </p>
          <p>You are never allowed to move down. The reason for this is to avoid the obvious strategy of moving down before a rebuild then moving back up. If a team is sold, the new owner can decide on either Tier 1 or 2 with no delay.</p>
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
      <v-expansion-panel-header class="grey--text">Tanking and Collusion</v-expansion-panel-header>
        <v-expansion-panel-content> <br/>
          <p>We define tanking as making decisions designed primarily to increase the chances of losing while not providing a substantiative benefit to your team beyond improving draft picks. Tanking is damaging to league morale, and it will not be tolerated. If the majority of the league believes a transaction is intentionally designed to tank, the commissioner can reject the transaction. Repeated issues will result in forfeiting of team.</p>
          <p>Collusion is defined as one team making a move that benefits another team without improving their own position. This is much more common in redraft leagues than dynasty, but it is possible. Examples that are more subtle include things like player “rentals” where you trade a player to someone competing and they later trade that player back as well as a draft pick. Transactions like this are not allowed. However, as noted above, trades intended to free up cap space are allowed.</p>
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
      <v-expansion-panel-header class="grey--text">Commissioner's Discretion</v-expansion-panel-header>
        <v-expansion-panel-content> <br/>
          <p>Due to the incredible complexity of this league, the commissioner needs some discretion to resolve issues. I will try to provide relief to teams in situations where a misunderstanding of the rules: </p>
          <ol type="1">
            <li>has no impact on the outcome of any matchups,</li>
            <li>hinders a team's ability to compete in either the short or long-term, and</li>
            <li>is brought to my attention promptly.</li>
          </ol>
          <p></p>
          <p>I will always strive to fairly resolve all situations. This league is meant to be fun. If you aren’t having fun, I will help you find a new owner for your team.</p>
        </v-expansion-panel-content>
      </v-expansion-panel>
    </v-expansion-panels>
    
  </div>
  </v-card>
</template>

<script>
export default {
  name: "RuleSummary.vue",
  data () {
    return {
      panel: []
    }
  }
}
</script>

<style scoped>

</style>