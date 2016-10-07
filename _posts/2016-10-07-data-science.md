---
layout: post
---

#### Data Science Book

Found a online data science book by Avrim Blum, John Hopcroft
and Ravindran Kannan, 440 pages.
[Foundation of data science](https://www.cs.cornell.edu/jeh/book2016June9.pdf)

#### Excel learning

Watched the presentation of
[Advanced Excel](https://www.youtube.com/watch?v=0nbkaYsR94c)
by Joel Spolsky, and learned few things.

- Use tables. You can make the whole relational model on one spreadsheet
by using tables. Add a new table by typing the column names, some data, then 
Data -> Insert Table, use "My table has headers". 

    1. With the tables you can rename the columns, referencing to the
       columns by names.
       In formulas, the current table is referred to as @ symbol.
       The columns in other tables are referred as TABLE[COLUMN].
       When you rename a column, all formulas are updated automatically.

    2. Add new records by pressing TAB on the last record in the last column.

    3. Include summaries by clicking Design -> Total row. You can change
       aggregate functions such as SUM and AVG.

    4. Under Design, change the style of a table, including highlighting
       for header and alternating colors for the rows.

- Instead of VLOOKUP, use a combination of MATCH and INDEX. They are faster. 
  The formula uses a location from the current table (symbol @) to find an index
  of a matching location from TaxRates table. Use the index to return a tax rate.

       =[@Salary]*INDEX(TaxRates[Tax Rate],MATCH([@Location], TaxRates[Location]))

  for the data model:

       Employees
           Salary
           Location

       TaxRates
           Location (PK)
           TaxRate

- Internally, Excel is using a relative reference to other columns (R1C1 format),
  so A10 is refered from B10 as RC[-1], or R[0]C[-1] - the same row, prior column. 

- Use pivot tables, they are an answer to any summarization reporting.

Few things I've learned about Excel myself today:

- You can change the style of a whole spreadsheet. 

- "Page Layout" -> Themes. I like Equity theme. 

- When you need to reference a static cell, add a name to it (field to the left
of cell formula on top.) Name it Taxrate, and then you can use Taxrate in
formulas instead of $B$5. 
Use Formulas -> Name Manager to edit, add, or delete names.



