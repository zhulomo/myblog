import { useEffect, useState } from "react";
import { data, Navigate, useLocation, useNavigate, useParams} from "react-router-dom"
import Nav from "../../components/nav";
import "./Update.css";

function Update() {
    //const { id } = useParams();
    const location  = useLocation();
    const data = location.state?.data;
    const [IsLogin, setIsLogin] = useState(false);
    const navigate = useNavigate();
    const [msg, setMsg] = useState("");
    const [formData, setFormData] = useState({
        id: 0,
        title: "",
        content: "",
        abstract: "",
    });
    //使用data?.article而不是data.article，更安全不会报错，报错显示undefined
    const article = data?.article
    console.log("接受的数据：",article.Id)
    const id =  article.Id;
    console.log("id是", id)
    useEffect(() => {
        if(localStorage.getItem("username")  ==  null) {
            navigate("/login");
        }
    });

    useEffect(() => {
        if (article) {
            setFormData({
                id: article.Id || "",
                title: article.Title || "",
                content: article.Content || "",
                abstract: article.Abstract || "",
            });
        }
    }, [article]);

    function handleChange(e) {
        const { name, value } = e.target;
        setFormData(prev => ({ ...prev, [name]: value}));
    }
    async function handleSubmit(e) {
        e.preventDefault();

        try {
            const res = await fetch(`http://localhost:8080/article/update?id=${id}`, {
                method: "POST",
                headers: {"Content-Type": "application/json"},
                body: JSON.stringify(formData),
                credentials: "include"
            })
            const data = await res.json();
            if (data.code === 1) {
                setMsg("修改成功")
                setTimeout(() => navigate("/article/list"), 2000);
            } else {
                setMsg(data.msg || "修改失败");
            }
        } catch (err) {
            console.error(err);
            setMsg("请求失败，请检查网络");
        }
    };

    return (
        <div className="article-page">
        <a href="/" className="back-link">← 返回首页</a>
        <h2>修改文章</h2>

            <form className="article-form" onSubmit={handleSubmit}>
                <div>
                <label htmlFor="title">标题</label>
                <input type="text" 
                value={formData.title || ""}
                id="title"
                name="title"
                onChange={handleChange}
                required />
                </div>
                <div>
                <label htmlFor="abstract">摘要</label>
                <input type="text" 
                id="abstract"
                name="abstract"
                value={formData.abstract || ""} 
                onChange={handleChange}
                required/>
                </div>

                <div>
                <label htmlFor="content">内容</label>
                <textarea  
                id="content"
                name="content"
                value={formData.content || ""}
                onChange={handleChange}
                ></textarea>
                </div>

                <button type="submit">修改</button>
            </form>
            {msg && <p className="message">{msg}</p>}
        </div>
  );
}

export default Update;