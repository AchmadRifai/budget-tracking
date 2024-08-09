import { createTheme, ThemeProvider } from "@mui/material/styles"

const theme = createTheme()

export default function DashboardLayout({ children }) {
    return <ThemeProvider theme={theme}></ThemeProvider>
}