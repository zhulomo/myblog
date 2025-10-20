import { useState } from "react";
import { useNavigate} from "react-router-dom"
import "./Register.css";

function Register() {
    const [username,  setUsername] = useState("");
    const [password, setPassword] =  useState("");
    const [repassword, setRepassword] = useState("");
    const [msg, setMsg] = useState("");
    const navigate = useNavigate();
    const handleSubmit = async (e) => {
        e.preventDefault();//阻止表单默认提交行为

        try {
            const res = await fetch("http://localhost:8080/register", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ 
                    username: username,
                    password: password,
                    repassword: repassword}),
            });

            const data = await res.json();

            if (data.code === 1) {
                setMsg("注册成功")
                setTimeout(() => navigate("/"), 2000);
            } else {
                setMsg(data.msg || "用户名已存在");
            }
        } catch (err) {
            console.error(err);
            setMsg("请求失败，请检查网络连接");
        }
    };

    return (

        
    <div className="register-container">
      <h2>注册</h2>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="用户名"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          required
        />
        <input
          type="password"
          placeholder="密码"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />
        <input
          type="password"
          placeholder="重复密码"
          value={repassword}
          onChange={(e) => setRepassword(e.target.value)}
          required
        />
        <button type="submit">注册</button>
      </form>

      {msg && <p className="message">{msg}</p>}
    </div>
  );
}

export default Register;





