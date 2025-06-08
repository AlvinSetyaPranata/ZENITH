import { useNavigate } from "@solidjs/router";

export default function Login() {
  const navigate = useNavigate();

  return (
    <div>
      <h1>Login Page Test</h1>
      <button onClick={() => alert("Clicked!")}>Test Button</button>
      <button onClick={() => navigate("/Dashboard")}>Go to Dashboard</button>
    </div>
  );
}
