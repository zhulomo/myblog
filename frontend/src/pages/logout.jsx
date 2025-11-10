import { useEffect } from "react";
import { useNavigate } from "react-router-dom";


function Logout() {
    const navigate = useNavigate();

    useEffect(() => {
        
        fetch("http://localhost:8080/api/logout", {
            method: "GET",
            credentials: "include",
        })
        .then(() => {
            alert("已退出登录");
            localStorage.removeItem("username");
            navigate("/login");
        })
        .catch(() => {
            alert("退出失败");
            navigate("/article/list");
        });
        
    },[navigate]);

    return <div>正在退出登录...</div>;
}
export default Logout;