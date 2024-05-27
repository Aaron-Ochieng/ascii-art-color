
## Ascii-art-Color
Ascii-art is a program which consists in receiving a string as an argument and outputting the string in a colored graphic representation using ASCII.

What we mean is we write the string received using a colored ASCII characters. You can define the color of your choice.


### Running the project
To run the project open bash terminal and run.

```bash
git clone https://learn.zone01kisumu.ke/git/aaochieng/ascii-art-color.git
cd ascii-art-color
```

```go
go run . --color=red "Hello\n" | cat -E
```
The above command will output the below ascii art with the whole string painted ```red```


The program only uses ```ANSII``` and ```RGB``` color format only. If ``hex`` or ```hsl``` values are used, error 
```bash 
Error:  color not found
exit status 1
``` 
will be returned.

To use RGB run 
```bash
go run . '--color=rgb(255, 0, 0)' Hello
```

### Error Handling

The program has error handling for the following scenarios:
If the standard.txt file is empty.
If there are any issues while reading the standard.txt file.
If there are any issues while handling newline characters in the input string.

#### Note

This program is a simple implementation of ASCII Art Color and might need to be adjusted based on your specific use case.

### Contribution
Ascii-Art-Color is an open source project and we welcome contributions.

If you want to contribute, fork and make changes.
### Contributors

<table>
<tr>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://learn.zone01kisumu.ke/git/aaochieng>
            <img src=https://learn.zone01kisumu.ke/git/avatars/8a1b24358854eb12998a07c269542193?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Aaron/>
            <br />
            <sub style="font-size:14px"><b>Aoron Ochieng</b></sub>
        </a>
    </td>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://learn.zone01kisumu.ke/git/ebarsula>
            <img src=https://learn.zone01kisumu.ke/git/avatars/fa966ef34b0ccdfe772414745aeee49f?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Emmanuel/>
            <br />
            <sub style="font-size:14px"><b>Emmanuel Barsulai</b></sub>
        </a>
    </td>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://learn.zone01kisumu.ke/git/abrakingoo>
            <img src=https://learn.zone01kisumu.ke/git/avatars/c307852c0cb9222c1ea2c71f98ff2d51?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Abraham/>
            <br />
            <sub style="font-size:14px"><b>Abraham kingoo</b></sub>
        </a>
    </td>
</tr>
</table>