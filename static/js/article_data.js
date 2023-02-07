let article_data=document.querySelector('.r-side').querySelectorAll('div')[0].querySelectorAll('span')[0]
let column_data=document.querySelector('.r-side').querySelectorAll('div')[0].querySelectorAll('span')[1]
let boil_data=document.querySelector('.r-side').querySelectorAll('div')[0].querySelectorAll('span')[2]

article_data.onclick=function(){
    location.href='/creator/data/content/article'
}
column_data.onclick=function(){
    location.href='/creator/data/content/column'
}
boil_data.onclick=function(){
    location.href='/creator/data/content/pin'
}


let data=document.querySelector('.article_data_detail').querySelectorAll('h1')


async function ask_for_article_data(){
    
    let res = await fetch('http://localhost:8080/creator/data/content/article',{
        method:'post',
        headers:{
            'Content-type':'application/json'
        },
       
    });
    let res2 = await res.json();
    let result=res2
    let a=result.all_article
    let b=result.article_view_number
    let c=result.like_number
    let d=result.comment_number
    let e=result.collection_number
    data[0].innerHTML=a
    data[1].innerHTML=b
   
    data[2].innerHTML=c
    data[3].innerHTML=d
    data[4].innerHTML=e

}
ask_for_article_data()