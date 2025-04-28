import { Route, Router } from "@solidjs/router";
import Home from "./pages/Home";
import RootLayout from "./pages/RootLayout";

function App() {
  return (
    <Router root={RootLayout}>
      <Route path="/" component={<Home />} />
    </Router>
  );
}

export default App;
