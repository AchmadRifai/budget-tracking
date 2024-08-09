import { useEffect, useState } from "react"
import { useSelector } from "react-redux"
import { getChart } from "../api/be"
import MyPieChart from "../components/MyPieChart"

export default function Charts() {
    const auth = useSelector(s => s.position.auth)
    const [loading, setLoading] = useState(true), [charts, setCharts] = useState({})
    useEffect(() => {
        setLoading(true)
        getChart(auth).then(r => setCharts(r.charts)).catch(console.log).finally(() => setLoading(false))
    }, [])
    return <>
        {Object.keys(charts).map(title => <MyPieChart title={title} data={charts[title]} />)}
    </>
}