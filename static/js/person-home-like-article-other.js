let list=document.querySelector('.list')
let changing=document.querySelector('.likes').querySelectorAll('span')[1]
let myid = new URL(location.href).searchParams.get("id")
changing.onclick=function(){
    location.href='/user/praise?id='+myid
}
async function ask_for_like_article(){
    let obj={
        user_id:myid
    }
    let res = await fetch('http://localhost:8080/user/likes',{
        method:'post',
        headers:{
            'Content-type':'application/json'
        },
        body:JSON.stringify(obj)
    });
    let res2 = await res.json();
    let result=res2.msg;
    for(let i=0;i<result.length;i++){
        let div=document.createElement('div')
        list.appendChild(div)
        let span1=document.createElement('span')
        let span2=document.createElement('span')
        let span3=document.createElement('span')
        let span4=document.createElement('span')
        let span5=document.createElement('span')
        let div1=document.createElement('div')
        let div2=document.createElement('div')
        let div3=document.createElement('div')
        div.appendChild(span1)
        div.appendChild(span2)
        div.appendChild(span3)
        div.appendChild(span4)
        div.appendChild(span5)
        div.appendChild(div1)
        div.appendChild(div2)
        div.appendChild(div3)
         

        span1.innerHTML=result[i].author
        span2.innerHTML=result[i].date
        span3.innerHTML=result[i].label
        span4.innerHTML=result[i].article_title
        span5.innerHTML=result[i].article_content
        div1.innerHTML='浏览'+result[i].view
        div2.innerHTML='点赞'+result[i].like
        div3.innerHTML='评论'+result[i].comment

        div.id=result[i].id
        div.onclick=function(e){
            location.href='/content?id='+e.target.id
        }
       
       
    }
    
}
ask_for_like_article()