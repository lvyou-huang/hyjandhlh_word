let article_list = document.querySelector('.article_list')
async function ask_for_article() {
    let obj={
        user_id:localStorage.getItem('myself_id')
    }
    let res = await fetch('http://localhost:8080/creator/content/article/essays',{
        method:'post',
        headers:{
            'Content-type':'application/json'
        },
        body:JSON.stringify(obj)
    });
    let res2 = await res.json();
    let result=res2.msg;
    for (var i = 0; i < result.length; i++) {
        let div = document.createElement('div');
        let h2 = document.createElement('h2');
        let span1 = document.createElement('span');
        let span2 = document.createElement('span');
        let span3 = document.createElement('span');
        let button=document.createElement('button')
        article_list.appendChild(div);
        div.appendChild(h2)
        div.appendChild(span1)
        div.appendChild(span2)
        div.appendChild(span3)
        div.appendChild(button)

        h2.innerHTML = result[i].article_title;
        span1.innerHTML = result[i].date
        span2.innerHTML ='点赞'+ result[i].like
        button.innerHTML='删除'
        span3.innerHTML ='收藏'+ result[i].collection
        div.id=result[i].id
        h2.id=result[i].id

        span1.id=result[i].id;
        span2.id=result[i].id;
        span3.id=result[i].id;

        h2.onclick=function(e){
            location.href='/content?id='+e.target.id
        }
        /*div.onclick=function(e){
            location.href='/web/winter-work/article/article1.html?id='+e.target.id
        }*/
        let d=result[i].id
        button.onclick=async function(){
            let obj={
                article_id:d
            }
            let res = await fetch('http://localhost:8080/deletearticle',{
                method:'post',
                headers:{
                    'Content-type':'application/json'
                },
                body:JSON.stringify(obj)
            });
            let res2 = await res.json();
            if(res2.msg=='删除成功'){
                div.remove()
            }
        }
    }
}
ask_for_article()