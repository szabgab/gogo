

Following https://developer.fyne.io/started/


```
$ go mod init github.com/szabgab/gogo/gui
```

created the go.mod file

```
$ go get fyne.io/fyne/v2

go: downloading fyne.io/fyne/v2 v2.1.2
go: downloading fyne.io/fyne v1.4.3
go get: added fyne.io/fyne/v2 v2.1.2

````


### TODO

* Download the content of a course
* From the current course page allow the user to go to the language selector page to switch to another language


### App starts
* If we already have a selected course then go to the course page
* If not selected course, but we alread have a list of courses then show the course selector
* If not courses then show the welcome splash screen and try to download list of languages and save it to our database and go to the course selector
* If we cannot download the list of courses in a reasonable time say so.

### Course selector
* List of the avilable courses
* Once the user selects a course save it in our database and go to the course page.
