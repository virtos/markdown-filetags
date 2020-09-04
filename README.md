# The question

I like to have control over my notes and keep them on my device. Should I use a directory structure or tags?

# The answer
You can have the best of both worlds! Create any folder structure you like and put tag names with a '+' sign.

# How does it work?
When you run this small utility, it looks for markdown files in the current folder and all subfolders, gathers all tags and creates an index file with all collected tags and reference to corresponding files. Just open _filetags.md file and you will see the power of markdown referencing.

# Do I need anything else?
A markdown editor. There are a number of great options. This is what I use:
    - Windows: [Typora](https://typora.io/)
    - Android: [Markor](https://github.com/gsantner/markor) or [Epsilon notes](http://epsilonexpert.com)

# Example
Suppose you decide to keep your notes in "My Notes" folder. You may put your web links into "Links" folder and your contacts into "Contats" folder. Here is a sample structure of your notes:

```
My Notes
    Contats
        Bob +friend +work.md
        Laura +work.md      
    Links                                 
        epsilon +editor.md              
        microsoft +work.md              
        typora +editor.md                
```
After running markdown-filetags.exe you will have index files added for each subfolder:
```
My Notes
    _filetags.md
    Contats
        _filetags.md
        Bob +friend +work.md
        Laura +work.md      
    Links                                 
        _filetags.md
        epsilon +editor.md              
        microsoft +work.md              
        typora +editor.md                
```
Contents of the _filetags.md file in "My Notes" folder:

----

<details markdown='1'><summary markdown='1'>contats</summary>
<li><a href="https://github.com/virtos/markdown-filetags/blob/master/README.md">Contats/Bob +friend +work</a>
<li><a href="https://github.com/virtos/markdown-filetags/blob/master/README.md">Contats/Laura +work</a>
</details>

<details markdown='1'><summary markdown='1'>editor</summary>
<li><a href="https://github.com/virtos/markdown-filetags/blob/master/README.md">Links/epsilon +editor</a>
<li><a href="https://github.com/virtos/markdown-filetags/blob/master/README.md">Links/typora +editor</a>
</details>

<details markdown='1'><summary markdown='1'>friend</summary>
<li><a href="https://github.com/virtos/markdown-filetags/blob/master/README.md">Contats/Bob +friend +work</a>
</details>

<details markdown='1'><summary markdown='1'>links</summary>
<li><a href="https://github.com/virtos/markdown-filetags/blob/master/README.md">Links/epsilon +editor</a>
<li><a href="https://github.com/virtos/markdown-filetags/blob/master/README.md">Links/microsoft +work</a>
<li><a href="https://github.com/virtos/markdown-filetags/blob/master/README.md">Links/typora +editor</a>
</details>

<details markdown='1'><summary markdown='1'>work</summary>
<li><a href="https://github.com/virtos/markdown-filetags/blob/master/README.md">Contats/Bob +friend +work</a>
<li><a href="https://github.com/virtos/markdown-filetags/blob/master/README.md">Contats/Laura +work</a>
<li><a href="https://github.com/virtos/markdown-filetags/blob/master/README.md">Links/microsoft +work</a>
</details>

----

As you can see the program collected tags for all notes. Clicking on a hyperlink will open a corresponding note.  

Index files in subfolders will only reference notes below them, so if you only want to see contacts related to work, go to contacts folder and find work tag in the index file.

Contents of the _filetags.md file in "Contacts" folder:

----
<details markdown='1'><summary markdown='1'>friend</summary>
<li><a href="https://github.com/virtos/markdown-filetags/blob/master/README.md">Bob +friend +work</a>
</details>

<details markdown='1'><summary markdown='1'>work</summary>
<li><a href="https://github.com/virtos/markdown-filetags/blob/master/README.md">Bob +friend +work</a>
<li><a href="https://github.com/virtos/markdown-filetags/blob/master/README.md">Laura +work</a>
</details>

----


# Unrelated notes
Using plain file structure with markdown notes is a very powerful system allowing you to keep your data private and secure. It's just text, so you will never lose it, you will never have to pay someone to keep them and you will never worry about your note application developer going out of business or doing something stupid. Enjoy!