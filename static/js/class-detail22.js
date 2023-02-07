let my_course_id = new URL(location.href).searchParams.get("id")

const select=(s)=>document.querySelector(s)
let class_head=select('.class_head')
let cover=document.querySelector('.class_head').querySelector('img')
let title=select('h1')
let abstract=class_head.querySelector('span')
let buy1=class_head.querySelectorAll('button')[0]
let trial=class_head.querySelectorAll('button')[1]

buy1.onclick =async function () {
    let obj={
        course_id:my_course_id
    }
    let res = await fetch('http://localhost:8080/purchase',{
        method:'post',
        headers:{
            'Content-type':'application/json'
        },
        body:JSON.stringify(obj)
    });
    let res2 = await res.json();
    if(res2.msg=='购买成功'){
        location.href='/book?id='+my_course_id
    }
}


let class_introduce=select('.class_introduce')
let intro=class_introduce.querySelector('div').querySelectorAll('span')[0]
let catelogue=class_introduce.querySelector('div').querySelectorAll('span')[1]
let introduce_content=select('.introduce_content')


async function ask_for_class(){
    let obj={
        course_id:my_course_id
    }
    let res = await fetch('http://localhost:8080/book',{
        method:'post',
        headers:{
            'Content-type':'application/json'
        },
        body:JSON.stringify(obj)
    });
    let res2 = await res.json();
    let result=res2.msg
    cover.src=result.cover
    title.innerHTML=result.course_title
    abstract.innerHTML=result.abstract
    if(res2.purchaseornot==0){
        buy1.innerHTML='去购买'
    }else{
        buy1.innerHTML='已购买'
    }
    introduce_content.innerHTML=result.introduction
    

}
ask_for_class()


catelogue.onclick=function(){
    location.href='/catelogue?id='+my_course_id
}

trial.onclick=function(){
    location.href='/catelogue?id='+my_course_id

}