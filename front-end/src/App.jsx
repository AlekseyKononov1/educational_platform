import { useState } from "react";

export default function App() {
  const [data, setData] = useState("");

  const fetchData = async () => {
    try {
      const res = await fetch("http://localhost:8080/hello");
      const text = await res.text();
      setData(text);
    } catch (err) {
      console.error("Fetch error:", err);
    }
  };

  return (
    <div>
      <h1>Fetch Example</h1>

      <button onClick={fetchData}>Fetch Data</button>

      {data && <pre>{data}</pre>}
    </div>
  );
}
