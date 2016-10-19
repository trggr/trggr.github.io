---
layout: post
---

#### Accounting

Accounting example with balance calculations:

    Active Accounts
        Cash
        Expense
        Equipment

    Passive Accounts
        Liability
        Income
        Revenue
        Equity
        Capital

    Initial owner's investment: Cash, Owne, 10000

    Acct   Balance Acct  Balances
    -----  ------------  -------------
    Owne   passive               10000
    Cash   active          10000
                           ----- -----
                           10000 10000

    Pays rent with cash:   Rent, Cash, 100

    Acct   Balance Acct  Balances
    -----  ------------  -------------
    Owne   passive               10000 
    Cash   active           9900
    Rent   active            100
                           ----- -----
                           10000 10000

    Receives cash for a sale: Cash, Income, 50

    Acct   Balance Acct  Balances
    -----  ------------  -------------
    Owne   passive               10000
    Cash   active           9950
    Rent   active            100
    Income passive                  50
                          ------ -----
                           10050 10050

    Buys equipment: Equipment, Cash, 5200 

    Acct   Balance Acct  Balances
    -----  ------------  -------------
    Owne   passive               10000
    Cash   active           4750
    Rent   active            100
    Income passive                  50
    Equip  active           5200
                          ------ -----
                           10050 10050

    Borrows a loan: Cash, Loan, 11000 

    Acct   Balance Acct  Balances
    -----  ------------  -------------
    Owne   passive               10000
    Cash   active          15750
    Rent   active            100
    Income passive                  50
    Equip  active           5200
    Loan   passive               11000
                          ------ -----
                           21050 21050

    Pays salaries: Salary, Cash, 5000

    Acct   Balance Acct  Balances
    -----  ------------  -------------
    Owne   passive               10000
    Cash   active          10750
    Rent   active            100
    Income passive                  50
    Equip  active           5200
    Loan   passive               11000
    Salary active           5000
                          ------ -----
                           21050 21050

Interestingly, I tried to implement it in Excel with tables.

    Accounts
        Account
        Balance Sheet
        Type
        Debits
        Credits
        Balance

    Postings
        Debit
        Credit
        Amount
        Description

and two pivot tables: Assets and Liabilities. On Assets, I only show the accounts
with Balance Sheet = Assets, on liabilities only the accounts with balance
sheet = Liabilities.

Formulas to use:

    Account[Debits]     =SUMIF(Postings[Debit], [@Account], Postings[Amount])           
    Account[Credits]    =SUMIF(Postings[Credit], [@Account], Postings[Amount])          
    Account[Balance]    =ABS([@Debits] -[@Credits])         

Interesting Excel's feature is "Trace Precedence" to show the arrows with
cell dependencies.

