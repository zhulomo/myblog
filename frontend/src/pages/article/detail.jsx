import { useEffect, useState } from "react";
import { data, Navigate, useNavigate, useParams} from "react-router-dom"
import Nav from "../../components/nav";
import "./Detail.css";

function Detail() {
    const [article, setArticle] = useState({});
    const { id } = useParams();
    const [IsLogin, setIsLogin] = useState(false);
    const navigate = useNavigate();
    useEffect(() => {
        fetch(`http://localhost:8080/article/detail/${id}`)
       
        .then((res) => res.json())
        .then((data) => {
            console.log(data);
            setArticle(data.article);
            
        })
        .catch((err) =>  console.error("获取文章失败:", err));

        fetch("http://localhost:8080/loginverify", {
            method: "GET",
            credentials: "include",
        })
        .then(res => res.json())
        .then(data => {
            console.log("logindata:",data)
            setIsLogin(data.isLogin);
            console.log("login:", data.isLogin)
        })
        .catch(err => console.error("登录状态获取失败:", err));
        }, [id]);
    

    const handleUpdate = () => {
        navigate(`/article/update?id=${article.Id}`);
    };
    const handleDelete = () => {
    if (!window.confirm("确定要删除这篇文章吗？")) return;
    fetch("http://localhost:8080/article/delete", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ id: article.Id }),
      credentials: "include",
    })
      .then(res => res.json())
      .then(data => {
        if (data.code === 1) {
          alert("删除成功！");
          navigate("/article/list");
        } else {
          alert("删除失败！");
        }
      })
      .catch(err => console.error("删除失败:", err));
  };
    return (
        
        <div className="article-detail">
            <Nav />
            <header>
                <h1>{article.Title}</h1>
                <div className="meta">
                    作者：{article.Author}
                    |发布日期：{article.CreateTime}
                    |浏览量：{article.Id}
                </div>
            </header>
            <article className="content">{article.Content}</article>
            {IsLogin && (
                <ul className="action-buttons">
                    <li>
                        <button onClick={handleUpdate}>修改文章</button>
                    </li>
                    <li>
                        <button onClick={handleDelete}>删除文章</button>
                    </li>
                </ul>
            )}    
        </div>
       
    );


}
export default Detail;
