import { useEffect, useState } from "react";
import { useNavigate} from "react-router-dom"
import Nav from "../../components/nav";
import "./List.css";

function List() {
    const [articles, setArticles] = useState([]);
    const [IsLogin, setIsLogin] = useState(false);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        //加载文章
        fetch("http://localhost:8080/api/articles", {
            method: "GET",
            headers: { "Content-Type": "application/json" },
            credentials: "include"
            
        })
        .then((res) => res.json())
        .then((data) => {
            console.log(data.IsLogin);
            setIsLogin(data.IsLogin || false);

            if(data.code === 1 && Array.isArray(data.data)) {
                setArticles(data.data);
            } else {
                setArticles([]);
            }
        })
        .catch((err) => {
            
            setArticles([]);
        })
        .finally(() => setLoading(false));
    }, []);
    if (loading) {
        return <p>加载中...</p>;
    }
    if (articles.length === 0) {
        return <p>加载文章失败，请刷新页面</p>;
    }
    return (
      <>
      <Nav />
    <div id="article-list">    
      {articles.map((article) => (
        <div key={article.Id} className="card article-card">
          <div className="card-body">
            <h5 className="card-title">{article.Title}</h5>
            <p className="card-text">{article.Content}</p>
            <a
              href={`/article/detail/${article.Id}`}
              className="btn btn-primary btn-sm"
            >
              阅读全文
            </a>
          </div>
        </div>
      ))}
    </div>
    </>
  );
   
     
}

export default List;
