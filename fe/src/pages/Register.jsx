import { useDispatch } from "react-redux"
import { setPosition } from "../store/positionSlice"
import { createTheme, ThemeProvider } from "@mui/material/styles"
import { Avatar, Box, Button, Container, CssBaseline, Grid, Link, TextField, Typography } from "@mui/material"
import Copyright from "../components/Copyright"
import LockOutlined from "@mui/icons-material/LockOutlined"

const theme = createTheme()

export default function Register() {
    const dispatch = useDispatch()
    return <ThemeProvider theme={theme}>
        <Container component="main" maxWidth="xs">
            <CssBaseline />
            <Box sx={{ marginTop: 8, display: 'flex', flexDirection: 'column', alignItems: 'center', }}>
                <Avatar sx={{ m: 1, bgcolor: 'secondary.main' }}><LockOutlined /></Avatar>
                <Typography component="h1" variant="h5">Registration</Typography>
                <Box component='form' noValidate sx={{ mt: 3 }}>
                    <Grid container spacing={2}>
                        <Grid item xs={12} sm={6}>
                            <TextField autoComplete="given-name" name="firstName" required fullWidth id="firstName" label="First Name" autoFocus />
                        </Grid>
                        <Grid item xs={12} sm={6}>
                            <TextField required fullWidth id="lastName" label="Last Name" name="lastName" autoComplete="family-name" />
                        </Grid>
                        <Grid item xs={12}>
                            <TextField required fullWidth id="username" label="Username" name="username" autoComplete="username" />
                        </Grid>
                        <Grid item xs={12}>
                            <TextField required fullWidth id="email" label="Email Address" name="email" autoComplete="email" />
                        </Grid>
                        <Grid item xs={12}>
                            <TextField required fullWidth name="password" label="Password" type="password" id="password" autoComplete="new-password" />
                        </Grid>
                    </Grid>
                    <Button type='submit' fullWidth variant='contained' sx={{ mt: 3, mb: 2 }}>Register</Button>
                    <Grid container justifyContent='flex-end'>
                        <Link onClick={() => dispatch(setPosition('login'))} href='#' variant='body2'>Log In</Link>
                    </Grid>
                </Box>
            </Box>
            <Copyright sx={{ mt: 5 }} />
        </Container>
    </ThemeProvider>
}