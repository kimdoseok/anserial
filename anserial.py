#!/usr/bin/env python
# -*- coding: UTF-8 -*-
#----------------------------------------------------------------------------
# Name:         Anserial
# Author:       Doseok Kim
# Copyright:    (c)Madang Business Solutions LLC
#----------------------------------------------------------------------------

class AlphaNum():
  def __init__(self):
    self.digits = "0123456789ABCDEF"
    self.dlen = len(self.digits)
    
  def getNext(self,anumber):      
    bnumber = anumber.upper()
    if not bnumber:
      bnumber = self.digits[0]
    blen = len(bnumber)
    overflow = False
    for i in range(blen):
      pos = -1-1*i
      adigit = bnumber[pos]
      try:
        idx = self.digits.index(adigit)
      except:
        return "0"
      if idx==len(self.digits)-1:
        overflow = True
      #print(blen+pos==blen-1)
      #print(pos,idx,overflow,bnumber)
      
      if blen+pos==blen-1:
        if idx==self.dlen-1:
          bnumber=bnumber[:pos]+self.digits[0]
          overflow = True
          continue
        else:
          bnumber=bnumber[:pos]+self.digits[idx+1]
          return bnumber
      if overflow:
        #print(bnumber,idx,pos,bnumber[:pos],self.digits[0],bnumber[pos+1:],)
        if idx==self.dlen-1:
          bnumber=bnumber[:pos]+self.digits[0]+bnumber[pos+1:]
          overflow = True
          continue
        else:
          bnumber=bnumber[:pos]+self.digits[idx+1]+bnumber[pos+1:]
          return bnumber
      else:
        return bnumber
    if overflow:
      bnumber = self.digits[1]+bnumber
    return bnumber
    
if __name__=="__main__":
  an = Anserial()
  anum = ""
  with open("output.txt","wb") as f:
    for i in range(10000):
      #print(anum)
      anum = an.getNext(anum)
      f.write(bytes("%s\n"%(anum,),"utf8"))
