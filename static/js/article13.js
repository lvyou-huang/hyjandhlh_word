let h1 = document.querySelector('.article_information').querySelector('h1');
let creator_name = document.querySelector('.article_information').querySelector('.creator_name');
let creator_photo = document.querySelector('.article_information').querySelector('.creator_photo');
let create_time = document.querySelector('.article_information').querySelector('.create_time');
let read_number = document.querySelector('.article_information').querySelector('.read_number');
let article_content = document.querySelector('.article_content');
let textarea = document.querySelector('.comment').querySelector('textarea');
let button = document.querySelector('.comment').querySelector('button');
let commentList = document.querySelector('.comment_list');
let like = document.querySelector('.like').querySelector('span');
let like_number = document.querySelector('.like').querySelector('div');
let comments = document.querySelector('.comments').querySelector('span');
let comments_number = document.querySelector('.comments').querySelector('div');
let collect = document.querySelector('.collect').querySelector('span');
let my_id = new URL(location.href).searchParams.get("id")
let category = document.querySelector('.category')
let num;

let date = new Date();
let year = date.getFullYear();
let month = date.getMonth() + 1;
let dates = date.getDate();
let hour = date.getHours();
let minite = date.getMinutes();
let second = date.getSeconds();
let time = year + '-' + month + '-' + dates + ' ' + hour + ':' + minite + ':' + second;



async function render() {
    let obj = {
        id: my_id
    }
    let res = await fetch('http://localhost:8080/content', {
        method: 'post',
        headers: {
            'Content-type': 'application/json'
        },
        body: JSON.stringify(obj)
    })
    let res2 = await res.json();
    let result = res2.msg
    h1.innerHTML = result.article_title;
    creator_name.innerHTML = result.author;
    creator_name.id = result.author_id;

    let obj2 = {
        user_id: result.author_id
    }
    let r = await fetch('http://localhost:8080/getuserinfo', {
        method: 'post',
        headers: {
            'Content-type': 'application/json'
        },
        body: JSON.stringify(obj2)
    });
    let r2 = await r.json();


    creator_photo.src = r2.msg.cover;
    category.innerHTML = result.category;
    create_time.innerHTML = result.date;
    read_number.innerHTML = '浏览量' + result.view;
    article_content.innerHTML = result.article_content;
    like_number.innerHTML = result.like > 0 ? result.like : 0;
    comments_number.innerHTML = result.comment
    if (res2.likeornot == 0) {
        like.innerHTML = ''
    } else {
        like.innerHTML = ''
    }


    if (res2.collectornot == 0) {
        collect.innerHTML = ''
    } else {
        collect.innerHTML = ''
    }

    num = result.like
    if (num == 0) {
        like_number.style.display = 'none'
    } else {
        like_number.innerHTML = num;
    }


    if (result.comment == 0) {
        comments_number.style.display = 'none'
    } else {
        comments_number.innerHTML = result.comment
    }

}
render();



creator_name.onclick = function (e) {
    location.href = '/user?id=' + e.target.id
}


async function ask_for_comments() {
    let obj = {
        id: my_id
    }
    let res = await fetch('http://localhost:8080/getcomment', {
        method: 'post',
        headers: {
            'Content-type': 'application/json'
        },
        body: JSON.stringify(obj)
    })
    let res2 = await res.json();
    let result = res2.msg;
    let length = res2.length;
    for (var i = 0; i < length; i++) {
        if (result[i].parentid == 0) {
            let content = result[i].content;
            let authorid = result[i].userid;
            let comment_id = result[i].id;
            let time = result[i].time
            let obj2 = {
                user_id: authorid
            }
            let r = await fetch('http://localhost:8080/getuserinfo', {
                method: 'post',
                headers: {
                    'Content-type': 'application/json'
                },
                body: JSON.stringify(obj2)
            })
            let r2 = await r.json();

            let author = r2.msg.name;
            let photo = r2.msg.cover;

            let faDiv = document.createElement('div');
            commentList.appendChild(faDiv);
            let img = document.createElement('img')
            let h4 = document.createElement('h4')
            let sonDiv1 = document.createElement('div')
            let sonDiv2 = document.createElement('div')
            let sonDiv3 = document.createElement('div')
            let text = document.createElement('textarea')
            let but = document.createElement('button')
            but.innerHTML = '发布'
            but.id = comment_id;
            faDiv.appendChild(img);
            faDiv.appendChild(h4)
            faDiv.appendChild(sonDiv1)
            faDiv.appendChild(sonDiv2)
            faDiv.appendChild(sonDiv3)
            faDiv.appendChild(text)
            faDiv.appendChild(but)
            img.src = photo;
            h4.innerHTML = author;
            sonDiv1.innerHTML = content;
            sonDiv2.innerHTML = time
            sonDiv3.innerHTML = '回复';
            sonDiv3.onmousemove = function () {
                sonDiv3.style.cursor = 'pointer'
            }
            sonDiv3.onclick = function () {
                if (faDiv.offsetHeight == 151) {
                    faDiv.style.height = '260px'
                } else {
                    faDiv.style.height = '150px'
                }
            }
            but.onclick = async function (e) {
                let date = new Date();
                let year = date.getFullYear();
                let month = date.getMonth() + 1;
                let dates = date.getDate();
                let hour = date.getHours();
                let minite = date.getMinutes();
                let second = date.getSeconds();
                let time = year + '-' + month + '-' + dates + ' ' + hour + ':' + minite + ':' + second;

                let obj = {
                    content: text.value,
                    projectid: my_id,
                    parentid: e.target.id,
                    time: time
                }
                let res = await fetch('http://localhost:8080/comment', {
                    method: 'post',
                    headers: {
                        'Content-type': 'application/json'
                    },
                    body: JSON.stringify(obj)
                })
                let result = await res.json();
                if (result.msg == "评论成功") {
                    text.innerHTML == ''
                    faDiv.style.height = '150px'
                }
            }
        } else {
            let content = result[i].content;
            let authorid = result[i].userid;
            let comment_id = result[i].id;
            let time = result[i].time;
            let obj3 = {
                parentid: result[i].parentid
            }
            let re = await fetch('http://localhost:8080/commentorinfo', {
                method: 'post',
                headers: {
                    'Content-type': 'application/json'
                },
                body: JSON.stringify(obj3)
            })
            let re2 = await re.json()
            let be_comment = re2.msg.name;

            let obj4 = {
                user_id: authorid
            }
            let r = await fetch('http://localhost:8080/getuserinfo', {
                method: 'post',
                headers: {
                    'Content-type': 'application/json'
                },
                body: JSON.stringify(obj4)
            })
            let r2 = await r.json();
            let author = r2.msg.name;
            let photo = r2.msg.cover;
            let faDiv_2 = document.createElement('span')
            let img_2 = document.createElement('img')
            let h4_2 = document.createElement('h4')
            let sonDiv1_2 = document.createElement('div')
            let sonDiv2_2 = document.createElement('div')
            let sonDiv3_2 = document.createElement('div')
            commentList.appendChild(faDiv_2);
            faDiv_2.appendChild(img_2);
            faDiv_2.appendChild(h4_2)
            faDiv_2.appendChild(sonDiv1_2)
            faDiv_2.appendChild(sonDiv2_2)
            faDiv_2.appendChild(sonDiv3_2)

            img_2.src = photo;
            h4_2.innerHTML = author;
            sonDiv1_2.innerHTML = content;
            sonDiv2_2.innerHTML = time
            sonDiv3_2.innerHTML = '回复' + be_comment

        }
    }
}
ask_for_comments();

button.onclick = async function () {
    let msg = textarea.value;
    if (msg == '') {
        alert('请输入内容')
        return;
    }
    else {
        let obj = {
            content: msg,
            projectid: my_id,
            parentid: 0,
            time: time
        }
        let res = await fetch('http://localhost:8080/comment', {
            method: 'post',
            headers: {
                'Content-type': 'application/json'
            },
            body: JSON.stringify(obj)
        })
        let result = await res.json();
        let d = result.id
        if (result.msg == "评论成功") {


            let faDiv = document.createElement('div');
            commentList.appendChild(faDiv);
            let img = document.createElement('img')
            let h4 = document.createElement('h4')
            let sonDiv1 = document.createElement('div')
            let sonDiv2 = document.createElement('div')
            let sonDiv3 = document.createElement('div')
            let text = document.createElement('textarea')
            let but = document.createElement('button')
            but.innerHTML = '发布'
            but.id = d;


            faDiv.appendChild(img);
            faDiv.appendChild(h4)
            faDiv.appendChild(sonDiv1)
            faDiv.appendChild(sonDiv2)
            faDiv.appendChild(sonDiv3)
            faDiv.appendChild(text)
            faDiv.appendChild(but)
            let obj = {
                user_id: localStorage.getItem('myself_id')
            }
            let res = await fetch('http://localhost:8080/getuserinfo', {
                method: 'post',
                headers: {
                    'Content-type': 'application/json'
                },
                body: JSON.stringify(obj)
            });
            let res2 = await res.json();
            let result = res2.msg
            img.src = result.cover
            h4.innerHTML = result.name
            sonDiv1.innerHTML = textarea.value
            textarea.value = ''
            let date = new Date();
            let year = date.getFullYear();
            let month = date.getMonth() + 1;
            let dates = date.getDate();
            let hour = date.getHours();
            let minite = date.getMinutes();
            let second = date.getSeconds();
            let time = year + '-' + month + '-' + dates + ' ' + hour + ':' + minite + ':' + second;
            sonDiv2.innerHTML = time
            sonDiv3.innerHTML = '回复'
            sonDiv3.onmousemove = function () {
                sonDiv3.style.cursor = 'pointer'
            }
            sonDiv3.onclick = function () {
                if (faDiv.offsetHeight == 151) {
                    faDiv.style.height = '260px'
                } else {
                    faDiv.style.height = '150px'
                }
            }
            but.onclick = async function (e) {
                let date = new Date();
                let year = date.getFullYear();
                let month = date.getMonth() + 1;
                let dates = date.getDate();
                let hour = date.getHours();
                let minite = date.getMinutes();
                let second = date.getSeconds();
                let time = year + '-' + month + '-' + dates + ' ' + hour + ':' + minite + ':' + second;

                let obj = {
                    content: text.value,
                    projectid: my_id,
                    parentid: d,
                    time: time
                }
                let res = await fetch('http://localhost:8080/comment', {
                    method: 'post',
                    headers: {
                        'Content-type': 'application/json'
                    },
                    body: JSON.stringify(obj)
                })
                let result = await res.json();
                if (result.msg == "评论成功") {
                    text.innerHTML == ''
                    faDiv.style.height = '150px'
                }
            }




        }


    }
}



like.onclick = async function () {
    let obj = {
        id: my_id
    }
    let res = await fetch('http://localhost:8080/like', {
        method: 'post',
        headers: {
            'Content-type': 'application/json'
        },
        body: JSON.stringify(obj)
    })
    let result = await res.json();


    if (result.msg == '点赞成功') {
        like.innerHTML = ''
        num++
        like_number.style.display = 'block'
        like_number.innerHTML = num;



    }
    else if (result.msg == '取消点赞') {
        like.innerHTML = ''
        num--
        if (num == 0) {
            like_number.style.display = 'none'
        } else {
            like_number.style.display = 'block'
            like_number.innerHTML = num;
        }
    }

}

if (num == 0) {
    like_number.style.display = 'none'
} else {
    like_number.innerHTML = num;
}


collect.onclick = async function () {
    let date = new Date();
    let year = date.getFullYear();
    let month = date.getMonth() + 1;
    let dates = date.getDate();
    let hour = date.getHours();
    let minite = date.getMinutes();
    let second = date.getSeconds();
    let time = year + '-' + month + '-' + dates + ' ' + hour + ':' + minite + ':' + second;
    let obj = {
        article_id: my_id,
        date: time
    }
    let res = await fetch('http://localhost:8080/collection', {
        method: 'post',
        headers: {
            'Content-type': 'application/json'
        },
        body: JSON.stringify(obj)
    })
    let result = await res.json();
    if (result.msg == '收藏成功') {
        collect.innerHTML = ''
    }
}