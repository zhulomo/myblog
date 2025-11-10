import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
//import Nav from "../../components/nav";
import "./Add.css";

function Add() {
    const [isLogin, setIsLogin] = useState(false);
    const [title,  setTitle] = useState("");
    const [author, setAuthor] = useState("");
    const [content, setContent] = useState("");
    const [abstract, setAbstract] = useState("");
    const [msg, setMsg] = useState("");
    const navigate = useNavigate();
    useEffect(() => {
      // fetch("http://localhost:8080/loginverify", {
      //   method: "GET",
      //   credentials: "include",
      // })
      // .then(res => res.json())
      // .then(data => {
      //   if (!data.isLogin) {
      //     navigate("/login");
      //   }
      // })
      console.log(localStorage.getItem("username"))
      if (localStorage.getItem("username") == null) {
        navigate("/login");
      }
    });
    const handleSubmit = async (e) => {
        e.preventDefault();//阻止表单默认提交行为
        const author = localStorage.getItem("username") || "匿名";


        try {
            const res = await fetch("http://localhost:8080/api/article/add", {
                method: "POST",
                headers: {"Content-Type": "application/json"},
                body: JSON.stringify({
                    title: title,
                    author: author,
                    content: content,
                    abstract: abstract,
                }),
                credentials: "include"
            })
            const data = await res.json();
            if (data.code === 1) {
                setMsg("发布成功")
                setTimeout(() => navigate("/article/list"), 2000);
            } else {
                setMsg(data.msg || "发布失败");
            }
        } catch (err) {
            console.error(err);
            setMsg("请求失败，请检查网络");
        }
    };

        return (
    <div className="article-page">
      <a href="/" className="back-link">← 返回首页</a>
      <h2>发布文章</h2>

      <form className="article-form" onSubmit={handleSubmit}>
        <div>
          <label htmlFor="title">标题</label>
          <input type="text" 
          value={title}
          id="title"
          name="title"
          placeholder="请输入文章标题" 
          onChange={(e) => setTitle(e.target.value)}
          required />
        </div>
        <div>
          <label htmlFor="abstract">摘要</label>
          <input type="text" 
          id="abstract"
          name="abstract"
          value={abstract} 
          placeholder="请输入文章摘要" 
          onChange={(e) => setAbstract(e.target.value)}
          required/>
        </div>

        <div>
          <label htmlFor="content">内容</label>
          <textarea  
          id="content"
          name="content"
          value={content}
          placeholder="请输入文章内容"
          onChange={(e) => setContent(e.target.value)}
          ></textarea>
        </div>

        <button type="submit">发布</button>
      </form>
    </div>
  );
}

export default Add;
