let lunbo = document.querySelector('.lunbo');
let imglist = document.querySelector('#imglist');
let pre = document.querySelector('.pre');
let nex = document.querySelector('.nex');
let icos = document.getElementById('icolist').getElementsByTagName('li');
let icolist = document.getElementById('icolist');
let left = 0;
let timer;
let n;
run();
function run() {
    if (left <= -4000) {
        left = 0;
    }
    let m = Math.floor(-left / 1000);
    imglist.style.marginLeft = left + 'px';
    n = (left % 1000 == 0) ? n = 2500 : n = 1;
    left = left - 20;
    ico_change(m);
    timer = setTimeout(run, n);
}
function change_position(n) {
    let x = -(n * 1000);
    imglist.style.marginLeft = x + 'px';
    left = x;
}
lunbo.onmouseover = function () {
    pre.style.display = 'block';
    nex.style.display = 'block';
}
lunbo.onmouseout = function () {
    pre.style.display = 'none';
    nex.style.display = 'none';
}
pre.onclick = function () {
    let prego = Math.floor(-left / 1000) - 1;
    if (prego == -1) {
        prego = 3;
    }
    change_position(prego)
}
nex.onclick = function () {
    let nexgo = Math.floor(-left / 1000) + 1;
    if (nexgo == 4) {
        nexgo = 0;
    }
    change_position(nexgo)
}
function ico_change(m) {
    for (let i = 0; i < icos.length; i++) {
        icos[i].style.backgroundColor = '';
    }
    if (m < icos.length) {
        icos[m].style.backgroundColor = 'white'
    }
}
for (let i = 0; i < icos.length; i++) {
    icos[i].onclick = function () {
        change_position(i);
    }
}
imglist.onmouseover = function () {
    clearTimeout(timer);
}
imglist.onmouseout = function () {
    run();
}


let data=document.querySelector('.sjgl').querySelectorAll('h1')
async function creatorHome1() {
    let res = await fetch('http://localhost:8080/creator/home', {
        method: 'post',
        headers: {
            'Content-type': 'application/json'
        },
    });
    let res2 = await res.json();
    let result=res2
    data[0].innerHTML=result.followers_number
    data[1].innerHTML=result.article_view_number
    data[2].innerHTML=result.like_number
    data[3].innerHTML=result.comment_number
    data[4].innerHTML=result.collection_number

}
creatorHome1()

