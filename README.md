# PasswordGenerator
Strong and secure random password generator.  Does not store any data, and is completely private.  Does not require an internet connection.

## Useage
To use, compile the source files and run the program with the `randpw` command. Output will be generated like so-

    C:\Users\username>randpw
    includedRunes=ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_-+={}[]|\:;"'<>,.?/`~
    randomSequences=
    bE,*c#5>x5X2GI*Dd;d.j_hRhA(d:lyDG{Cosa~(aJyrsdo}taz(w$?Y'nPA"$WC=+$V'"oVt],/I4>tK%WY;;4Fw7L>H$+~kZwV
    `H{m$5^tRc_ez8x%9kiU;NXqfEN/JRg.g-pgGNF`2Q8LXq4@CEn_jJjU.b)nms_dCG;Y!V#s=7/9Kd|XOc<`|G[H0uNdCn!Ou[*4
    *:D(8m-\rU/>sBAd[Rgy:i*lR.znNalxv_*0DkeT:?f=+KHkDV6O%NgQsMeI#a2Dv{D9f4lEv+6|*[j4Z)xh66QpzyB1~k:8oKs#
    A'[Ny0}=T@Xsa-#^b~)`4:7Q.,HoGk*"wyhczAx,7?n'$.eWA#b(qkin%WN$mD3q.[wAQIs`HH1}6|eMd1m07gjZp($gl)y*X)N1
    GB&g+UA~FRMjeqe~5q*-yL<t05}<yjnl*5&'BO+Hm/qTYU""D\bs*_7{NEy%T4M10+gc}^CE<tJV\rJvR~+ctmZ[|Np|Wg\u\ohk

In this output, `includedRunes` is the list of all characters used for the password generation.  `randomSequences` consists of 5 random sequences of 100 characters.  From these sequences, you can choose any subsequence of any length for use as a password.

To exclude characters from the `includedRunes`, pass the list of characters you would like to exclude into the `randpw` command like so-

    C:\Users\username>randpw "!@#$%^&*()_-+={}[]|\:;""'<>,.?/`~"
    excludedRunes=!@#$%^&*()_-+={}[]|\:;"'<>,.?/`~
    includedRunes=ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789
    randomSequences=
    lUUpbR0lImOzPWMBra0j62tUS6UwyKD37ZfanHjFNsz5bApqj5Wb0jlVNMvQcU84eVDpYvDdbHuq4EfeYot8lcd4Y5ZuWtWFydy9
    6OjrRBZuQjcyC4tHPMVOenkF2HVq8FM5s6kkGUasA36n9k4P0xx3pj1si3pcQmWl69sURDEOPpMUJTekegjQed4qedwStv3mO6W2
    9bVhilwIsddKHyXa3oM6qaQUFRYTyLGAUrf6IXjoenmKdbHkEocarocQWCCjj6wcRnKUeg98yALE7dnrnjrYYbXZW9VHuvESMNT6
    2V2ziirzwx6x3qzNDsgFMSIE1vakj3hYoJlOtJL1jadxQDnOvw4nGtHwwi38rVNzdVOIc03Y3OzkSMbERidHzzVUIP5DdFNAMTSX
    UzJ5dkroOMogZoa6pt5lOfmPrZVP32f9WPwAkK72B7VOU9fIVdDbpXXjuZmXkuKLjQYnJoTqWZKtB5AHgEXel4i2YQoxHeQ76U5M

Note the following with this second example-
- Certain special characters (ex., `"`) will need to be properly escaped, to ensure that they are used verbatim within the program.  Keep in mind, certain special characters may have special meanings both to the shell, and also to the Go compiler.
- The `excludedRunes` output will display the full list of all characters to be excluded, as confirmation of your specified input.  These characters will not be included in the `randomSequences` output.
- The `includedRunes` output will display the full list of all characters to be included, as confirmation of your specified input.  The `excludedRunes` are filtered out of this list.
