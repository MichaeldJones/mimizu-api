var url = ""

function index(data){

  document.getElementById("books").innerHTML='';

  for (var i = 0; i < data.length; i++) {
    var book = document.createElement("a");
    book.href = data[i].link;
    book.className = "flex hover:bg-gray-200 m-2 border-b-2 border-r-2 border-l-2 border-gray-200 pl-4";


    var title = document.createElement("p");
    title.className = "w-1/3"
    var author = document.createElement("p");
    author.className = "w-1/3"
    var total  = document.createElement("p");
    total.className = "w-1/6"
    var unique = document.createElement("p");
    unique.className = "w-1/6"
    
    title.innerHTML = data[i].title;
    author.innerHTML = data[i].author;
    total.innerHTML = data[i].total;
    unique.innerHTML = data[i].unique;
    

    book.appendChild(title);
    book.appendChild(author);
    book.appendChild(total);
    book.appendChild(unique);

    document.getElementById("books").appendChild(book);
  }
}

function byTitle(){

  if(titleBool) {

    titleBool  = false;
    authorBool = true;
    totalBool  = true;
    uniqBool   = true;

    fetch(url+'/api/books/top/title')
    .then(function (res) {
      return res.json();
    })
    .then(function (data){
      index(data);
    })
    .catch(function (err){
      console.log('error: ' + err);
    })
  } else {

    titleBool = true;

    fetch(url+'/api/books/low/title')
    .then(function (res) {
      return res.json();
    })
    .then(function (data){
      index(data);
    })
    .catch(function (err){
      console.log('error: ' + err);
    })
  }

}

function byAuthor(){

  if(authorBool) {

    authorBool = false;
    titleBool  = true;
    totalBool  = true;
    uniqBool   = true;

    fetch(url+'/books/top/author')
    .then(function (res) {
      return res.json();
    })
    .then(function (data){
      index(data);
    })
    .catch(function (err){
      console.log('error: ' + err);
    })
  } else {

    authorBool = true;

    fetch(url+'/api/books/low/author')
    .then(function (res) {
      return res.json();
    })
    .then(function (data){
      index(data);
    })
    .catch(function (err){
      console.log('error: ' + err);
    })
  }

}

function byTotal(){

  if(totalBool) {

    totalBool  = false;
    titleBool  = true;
    authorBool = true;
    uniqBool   = true;

    fetch(url+'/api/books/top/total')
    .then(function (res) {
      return res.json();
    })
    .then(function (data){
      index(data);
    })
    .catch(function (err){
      console.log('error: ' + err);
    })
  } else {

    totalBool = true;

    fetch(url+'/api/books/low/total')
    .then(function (res) {
      return res.json();
    })
    .then(function (data){
      index(data);
    })
    .catch(function (err){
      console.log('error: ' + err);
    })
  }

}

function byUnique(){

  if(uniqBool) {

    uniqBool   = false;
    titleBool  = true;
    authorBool = true;
    totalBool  = true;

    fetch(url+'/api/books/top/uniq')
    .then(function (res) {
      return res.json();
    })
    .then(function (data){
      index(data);
    })
    .catch(function (err){
      console.log('error: ' + err);
    })
  } else {

    uniqBool = true;

    fetch(url+'/api/books/low/uniq')
    .then(function (res) {
      return res.json();
    })
    .then(function (data){
      index(data);
    })
    .catch(function (err){
      console.log('error: ' + err);
    })
  }
}

document.getElementById("search").addEventListener("keyup", function(e) {

  if(e.key == "Enter"){

    var inquery = document.getElementById("search").value;
    console.log(inquery);

    fetch(url+'/api/books/search/' + inquery)
    .then(function (res) {
      return res.json();
    })
    .then(function (data){
      index(data);
    })
    .catch(function (err){
      console.log('error: ' + err);
    })
  }
});

var titleBool  = true;
var authorBool = true;
var totalBool  = true;
var uniqBool   = true;

fetch(url+'/api/books')
.then(function (res) {
  return res.json();
})
.then(function (data){
  index(data);
})
.catch(function (err){
  console.log('error: ' + err);
})