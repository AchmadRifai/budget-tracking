import { Button } from "@mui/material"
import { useState } from "react"
import { useDispatch } from "react-redux"
import { dashboardLogout } from "../api/be"
import { setAuth, setPosition } from "../store/positionSlice"

export default function Dashboard() {
    const dispatch = useDispatch()
    const [loading, setLoading] = useState(false)
    const loggingOut = () => {
        setLoading(true)
        dashboardLogout().then(r => {
            localStorage.clear()
            dispatch(setPosition('login'))
            dispatch(setAuth(''))
        }).catch(e => {
            console.log(e)
            localStorage.clear()
            dispatch(setPosition('login'))
            dispatch(setAuth(''))
        }).finally(() => setLoading(false))
    }
    return <Button onClick={() => loggingOut()} disabled={loading}>Logout</Button>
}