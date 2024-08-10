import DashboardLayout from "../layouts/Dashboard"
import Users from "../features/Users"
import { useSelector } from "react-redux"
import Charts from "../features/Charts"
import Budgets from "../features/Budgets"
import Categories from "../features/Categories"
import Expenses from "../features/Expenses"

const features = [
    <Charts />,
    <Budgets />,
    <Categories />,
    <Expenses />,
    <Users />
]

export default function Dashboard() {
    const menu = useSelector(s => s.position.menu)
    return <DashboardLayout>
        {features[menu]}
    </DashboardLayout>
}