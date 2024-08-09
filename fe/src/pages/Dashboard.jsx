import DashboardLayout from "../layouts/Dashboard"
import Users from "../features/Users"
import { useSelector } from "react-redux"

const features = [
    <></>,
    <></>,
    <></>,
    <></>,
    <Users />
]

export default function Dashboard() {
    const menu = useSelector(s => s.position.menu)
    return <DashboardLayout>
        {features[menu]}
    </DashboardLayout>
}