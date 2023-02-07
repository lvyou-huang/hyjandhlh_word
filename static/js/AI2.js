let li = document.querySelector('nav').querySelector('ul').querySelectorAll('li');

let passage_content = document.querySelector('#passage_content');


li[0].onclick = function () {
    location.href = '../home'
}
li[1].onclick = function () {
    location.href = '../following'
}
li[2].onclick = function () {
    location.href = '../backend'
}
li[3].onclick = function () {
    location.href = '../frontend'
}
li[4].onclick = function () {
    location.href = '../android'
}
li[5].onclick = function () {
    location.href = '../ios'
}
li[6].onclick = function () {
    location.href = '../ai'
}
li[7].onclick = function () {
    location.href = '../freebie'
}
li[8].onclick = function () {
    location.href = '../career'
}
li[9].onclick = function () {
    location.href = '../read'
}



let common_list=document.querySelector('.common_list')
let new_list=document.querySelector('.new_list')
let hot_list=document.querySelector('.hot_list')


async function ask_for_data() {
    let res = await fetch('http://localhost:8080/ai',{
        method:'post',
        headers:{
            'Content-type':'application/json'
        }
    });
    let res2 = await res.json();
    let result=res2.msg;
    for (var i = 0; i < result.length; i++) {
        let div = document.createElement('div');
        passage_content.appendChild(div);
        let span1 = document.createElement('span');
        let span2 = document.createElement('span');
        let span3 = document.createElement('span');
        let h2 = document.createElement('h2');
        let img = document.createElement('img');
        div.appendChild(span1)
        div.appendChild(span2)
        div.appendChild(span3)
        div.appendChild(h2)
        div.appendChild(img)
        span1.innerHTML=result[i].author;
        span2.innerHTML=result[i].date;
        span3.innerHTML=result[i].category;
        h2.innerHTML=result[i].article_title;
        img.src=result[i].cover;
        div.id=result[i].id;
        h2.id=result[i].id;
        span1.id=result[i].id;
        span2.id=result[i].id;
        span3.id=result[i].id;
        img,id=result[i].id


        div.onclick=function(e){
            location.href='../content?id='+e.target.id
        }

    }

}
ask_for_data();




new_list.onclick=async function(){
    let t=passage_content.children.length;
    for(let i=0;i<t;i++){
        passage_content.removeChild(passage_content.children[0])
    }
    common_list.style.color='rgb(124, 121, 121)'
    hot_list.style.color='rgb(124, 121, 121)'
    new_list.style.color='rgb(79, 174, 237)'
    let res = await fetch('http://localhost:8080/ai?status=new',{
        method:'post',
        headers:{
            'Content-type':'application/json'
        }
    });
    let res2 = await res.json();
    let result=res2.msg;
    for (var i = 0; i < result.length; i++) {
        let div = document.createElement('div');
        passage_content.appendChild(div);
        let span1 = document.createElement('span');
        let span2 = document.createElement('span');
        let span3 = document.createElement('span');
        let h2 = document.createElement('h2');
        let img = document.createElement('img');
        div.appendChild(span1)
        div.appendChild(span2)
        div.appendChild(span3)
        div.appendChild(h2)
        div.appendChild(img)
        span1.innerHTML=result[i].author;
        span2.innerHTML=result[i].date;
        span3.innerHTML=result[i].category;
        h2.innerHTML=result[i].article_title;
        img.src=result[i].cover;
        div.id=result[i].id;
        h2.id=result[i].id;
        div.onclick=function(e){
            location.href='../content?id='+e.target.id
        }
    }
}

hot_list.onclick=async function(){
    let t=passage_content.children.length;
    for(let i=0;i<t;i++){
        passage_content.removeChild(passage_content.children[0])
    }
    common_list.style.color='rgb(124, 121, 121)'
    new_list.style.color='rgb(124, 121, 121)'
    hot_list.style.color='rgb(79, 174, 237)'
    let res = await fetch('http://localhost:8080/ai?status=hot',{
        method:'post',
        headers:{
            'Content-type':'application/json'
        }
    });
    let res2 = await res.json();
    let result=res2.msg;
    for (var i = 0; i < result.length; i++) {
        let div = document.createElement('div');
        passage_content.appendChild(div);
        let span1 = document.createElement('span');
        let span2 = document.createElement('span');
        let span3 = document.createElement('span');
        let h2 = document.createElement('h2');
        let img = document.createElement('img');
        div.appendChild(span1)
        div.appendChild(span2)
        div.appendChild(span3)
        div.appendChild(h2)
        div.appendChild(img)
        span1.innerHTML=result[i].author;
        span2.innerHTML=result[i].date;
        span3.innerHTML=result[i].category;
        h2.innerHTML=result[i].article_title;
        img.src=result[i].cover;
        div.id=result[i].id;
        h2.id=result[i].id;
        div.onclick=function(e){
            location.href='../content?id='+e.target.id
        }
    }
}


common_list.onclick=async function(){
    let t=passage_content.children.length;
    for(let i=0;i<t;i++){
        passage_content.removeChild(passage_content.children[0])
    }
    hot_list.style.color='rgb(124, 121, 121)'
    new_list.style.color='rgb(124, 121, 121)'
    common_list.style.color='rgb(79, 174, 237)'
    let res = await fetch('http://localhost:8080/ai',{
        method:'post',
        headers:{
            'Content-type':'application/json'
        }
    });
    let res2 = await res.json();
    let result=res2.msg;
    for (var i = 0; i < result.length; i++) {
        let div = document.createElement('div');
        passage_content.appendChild(div);
        let span1 = document.createElement('span');
        let span2 = document.createElement('span');
        let span3 = document.createElement('span');
        let h2 = document.createElement('h2');
        let img = document.createElement('img');
        div.appendChild(span1)
        div.appendChild(span2)
        div.appendChild(span3)
        div.appendChild(h2)
        div.appendChild(img)
        span1.innerHTML=result[i].author;
        span2.innerHTML=result[i].date;
        span3.innerHTML=result[i].category;
        h2.innerHTML=result[i].article_title;
        img.src=result[i].cover;
        div.id=result[i].id;
        h2.id=result[i].id;
        div.onclick=function(e){
            location.href='../content?id='+e.target.id
        }
    }
}

