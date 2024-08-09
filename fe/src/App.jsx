import { useSelector } from "react-redux"
import Login from "./pages/Login"
import Register from "./pages/Register"

function App() {
  const position = useSelector(state => state.position.position)
  const auth = useSelector(state => state.position.auth)
  console.log('position ', position, ' auth ', auth)
  return (
    <>
      {
        auth === '' ?
          position === 'login' ? <Login /> : <Register /> :
          <div></div>
      }
    </>
  )
}

export default App
