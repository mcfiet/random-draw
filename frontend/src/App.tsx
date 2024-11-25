import { useEffect, useState } from "react";
import { jwtDecode } from "jwt-decode";
import "./App.css";

interface UserData {
  username: string;
  password: string;
}

function App() {
  const url = "http://localhost:3000";
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [token, setToken] = useState<string | null>(
    localStorage.getItem("token"),
  );
  const [draw, setDraw] = useState<string | null>(null);

  const getDrawQuery = async () => {
    try {
      const response = await fetch(url + "/draw", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: "Bearer " + token!,
        },
      });
      console.log(response);
      if (!response.ok) {
        throw new Error("Failed to get draw");
      }
      const resDraw = await response.text();
      console.log("Draw: ", resDraw);
      setDraw(resDraw);
      setError(null);
    } catch (error) {
      setError(error.message);
    } finally {
      setLoading(false);
      setError(null);
    }
  };

  useEffect(() => {
    if (token) {
      localStorage.setItem("token", token);
      getDrawQuery();
    }
  }, [token]);

  const loginQuery = async (loginData: UserData) => {
    try {
      const response = await fetch(url + "/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(loginData),
      });
      if (!response.ok) {
        throw new Error("Failed to login");
      }
      const jwtToken = response.headers.get("Authorization")?.split(" ")[1];
      setToken(jwtToken ?? "");
      setError(null);
    } catch (error) {
      setError(error.message);
    } finally {
      setLoading(false);
      setError(null);
    }
  };

  function handleLogin(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();
    const formData = new FormData(event.currentTarget);
    const username = formData.get("username") as string;
    const password = formData.get("password") as string;
    const loginData = { username, password };
    loginQuery(loginData);
  }

  async function drawHandler() {
    console.log(url + "/draw");
    try {
      const response = await fetch(url + "/draw", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: "Bearer " + token!,
        },
      });
      if (!response.ok) {
        throw new Error(await response.text());
      }
      const reciever = await response.text();
      console.log("drawn user: ", reciever);
      setDraw(reciever);
      setError(null);
    } catch (error) {
      setError(error.message);
    } finally {
      setLoading(false);
    }
  }

  function logoutHandler() {
    localStorage.clear();
    window.location.href = "/";
  }

  if (loading) {
    return <h1>Loading...</h1>;
  }

  // if (error) {
  //   return <h1>{error}</h1>;
  // }

  if (!token) {
    return (
      <>
        <h1>Anmeldung</h1>
        <form onSubmit={handleLogin}>
          <input type="text" name="username" placeholder="Username" />
          <input type="password" name="password" placeholder="Password" />
          <button type="submit">Login</button>
        </form>
        <div>{token}</div>
      </>
    );
  }

  return (
    <div>
      <h1>Hallo {jwtDecode(token).username}</h1>
      {draw && (
        <div style={{ margin: "1rem" }}>{`Du hast ${draw} gezogen`}</div>
      )}
      <button style={{ marginBottom: "1rem" }} onClick={drawHandler}>
        Auslosen
      </button>
      {error && (
        <>
          <div style={{ color: "lightpink" }}>{error}</div>
        </>
      )}
      <div>
        <button onClick={logoutHandler}>Logout</button>
      </div>
    </div>
  );
}

export default App;
