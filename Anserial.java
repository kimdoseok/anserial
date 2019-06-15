package com.madang.anserial;

public class Anserial {
	static String digits = "0123456789ABCDEF";
	static String prefix = "";

	public Anserial() {
		// TODO Auto-generated constructor stub
	}
	
	public String getNext(String anum) {

		int dlen = digits.length();
		int plen = prefix.length();
		String bnum = anum.substring(plen,anum.length());
		int blen = bnum.length();
		
		boolean validnum = true;
		
		for (char x:bnum.toCharArray()) {
			if (digits.indexOf(x)<0) {
				validnum = false;
				break;
			}
		}
		if (!validnum) {
			return prefix+digits.toCharArray()[0];
		}
		
		boolean zero = false;
		if (digits.toCharArray()[0] == '0') {
			zero = true;
		}
		boolean overflow = false;
		
		for (int i=0;i<blen;i++) {
			int pos = blen - 1 - 1*i; // index of given string
			char adigit = bnum.toCharArray()[pos];
			int idx = digits.indexOf(adigit); // index of all digits
			if (idx < 0) {
				return prefix + digits.toCharArray()[0];
			}
			if (idx == dlen-1) {
				overflow = true;
			}

			if (pos == blen-1) {
				if (idx == dlen-1) {
					bnum = bnum.substring(0,pos) + digits.toCharArray()[0];
					overflow = true;
					continue;
				} else {
					bnum = bnum.substring(0,pos) + digits.toCharArray()[idx+1];
					return prefix + bnum;

				}
			}
			if (overflow) {
				if (idx == dlen-1) {
					bnum = bnum.substring(0,pos) + digits.toCharArray()[0] + bnum.substring(pos+1,blen);
					overflow = true;
					continue;
				} else {
					bnum = bnum.substring(0,pos) + digits.toCharArray()[idx+1] + bnum.substring(pos+1,blen);
					return prefix + bnum;
				}
			} else {
				return prefix + bnum;
			}
		}

		if (overflow) {
			if (zero) {
				bnum = digits.toCharArray()[1] + bnum;
			} else {
				bnum = digits.toCharArray()[0] + bnum;
			}
		}

		return prefix + bnum;
		
	}

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Anserial ans = new Anserial();
		String num = prefix + digits.charAt(0);
		for (int i=0;i<100000;i++) {
			System.out.println(num);
			num = ans.getNext(num);			
		}
	}
}
