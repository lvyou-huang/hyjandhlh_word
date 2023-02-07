let lunbo = document.querySelector('.lunbo');
let imglist = document.querySelector('#imglist');
let pre = document.querySelector('.pre');
let nex = document.querySelector('.nex');
let icos = document.getElementById('icolist').getElementsByTagName('li');
let icolist = document.getElementById('icolist')
let left = 0;
let timer;
run();
function run() {
    if (left <= -4000) {
        left = 0;
    }
    let m = Math.floor(-left / 1000);
    imglist.style.marginLeft = left + 'px';
    var n = (left % 1000 == 0) ? n = 1500 : n = 5;
    left = left - 10;
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
for(let i=0;i<icos.length;i++){
    icos[i].onclick=function(){
        change_position(i);
    }
}