let li1 = document.querySelector('#nav1').querySelector('ul').querySelectorAll('li');
let li2 = document.querySelector('#nav2').querySelector('ul').querySelectorAll('li');
li1[0].onclick = function () {
    location.href = '/course';
}
li1[1].onclick = function () {
    location.href = '/course/backend';
}
li1[2].onclick = function () {
    location.href = '/course/frontend';
}
li1[3].onclick = function () {
    location.href = '/course/android';
}
li1[4].onclick = function () {
    location.href = '/course/ios';
}
li1[5].onclick = function () {
    location.href = '/course/ai';
}
li1[6].onclick = function () {
    location.href = '/course/tool';
}
li1[7].onclick = function () {
    location.href = '/course/code';
}
li1[8].onclick = function () {
    location.href = '/course/read';
}


