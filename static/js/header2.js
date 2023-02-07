let h_li=document.querySelector('.header').querySelectorAll('li');
let h_search=document.querySelector('input');
let h_search_logo=document.querySelector('.search_logo');
let h_create=document.querySelector('.header_create');
let h_person=document.querySelector('.header_person');
h_li[0].onclick=function(){
    location.href='/home'
}
h_li[2].onclick=function(){
    location.href='/course'
}
h_li[1].onclick=function(){
    alert('暂无相关内容');
}
for(var i=3;i<=8;i++)
{
    h_li[i].onclick=function(){
    alert('暂无相关内容');
    }
}

h_create.onclick=function(){
    location.href='/creator/home';
}
h_person.onclick=function(){
    location.href='/user?id='+localStorage.getItem('myself_id');
}
async function ask_for_photo(){
    let obj={
        user_id:localStorage.getItem('myself_id')
    }
    let res = await fetch('http://localhost:8080/getuserinfo',{
        method:'post',
        headers:{
            'Content-type':'application/json'
        },
        body:JSON.stringify(obj)
    });
    let res2 = await res.json();
    h_person.src=res2.msg.cover

}
ask_for_photo()


let h_search_content=document.querySelector('.header').querySelector('input')
h_search_logo.onclick=function(){
        location.href='/search?content='+ h_search_content.value
        h_search_content.innerHTML=' '
}