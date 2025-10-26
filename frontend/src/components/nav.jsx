import { Link } from "react-router-dom"
import { useEffect, useState } from "react"
import "./Nav.css"

function Nav() {
    const [isLogin, setIsLogin] = useState(false);
    const [username, setUsername] = useState("");

    useEffect(() => {
        fetch("http://localhost:8080/loginverify", {
            method: "GET",
            credentials: "include",
        })
        .then(res => res.json())
        .then(data => {
          console.log("data1",data)
            if (data.isLogin) {
                console.log(data.isLogin)
                setIsLogin(true);
                setUsername(data.username);
            }
             else {
                setIsLogin(false);
             }
        })
        .catch(err => console.error("登录状态获取失败：", err));
    }, []);

    return (
    <nav className="navbar">
      <div className="navbar-container">
        <Link className="navbar-brand" to="/">
          MyBlog
        </Link>
        <ul className="nav-list">
          <li className="nav-item">
            <Link className="nav-link" to="/">
              首页
            </Link>
          </li>
          {!isLogin && (
            <>
              <li className="nav-item">
                <Link className="nav-link" to="/login">
                  登录
                </Link>
              </li>
              <li className="nav-item">
                <Link className="nav-link" to="/register">
                  注册
                </Link>
              </li>
            </>
          )}
          {isLogin && (
            <>
              <li className="nav-item">
                <Link className="nav-link" to="/article/add">
                  发布文章
                </Link>
              </li>
              <li className="nav-item">
                <Link className="nav-link" to="/logout">
                  退出
                </Link>
              </li>
              <li className="nav-item">
                <Link className="nav-link" to="/profile">
                  个人中心
                </Link>
              </li>
            </>
          )}
        </ul>
      </div>
    </nav>
  );
}


export default Nav;