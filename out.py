import sqlite3
from xlsxwriter.workbook import Workbook

workbook = Workbook('output.xlsx')
worksheet = workbook.add_worksheet()
# 传入数据库路径，db.s3db或者test.sqlite
conn=sqlite3.connect('ask.db')
c=conn.cursor()
mysel=c.execute("select * from tables")
for i, row in enumerate(mysel):
    for j, value in enumerate(row):
        worksheet.write(i, j, value)
workbook.close()
