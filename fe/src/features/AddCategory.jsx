import { Add } from "@mui/icons-material"
import { Button, Dialog, DialogActions, DialogContent, DialogTitle, Fab, TextField } from "@mui/material"
import { useState } from "react"
import { addCategory } from "../api/be"
import { useSelector } from "react-redux"

export default function AddCategory({ onClose }) {
    const auth = useSelector(s => s.position.auth)
    const [open, setOpen] = useState(false), [name, setName] = useState(''), [loading, setLoading] = useState(false)
    const closing = () => {
        setOpen(false)
        onClose()
    }
    const saving = () => {
        setLoading(true)
        addCategory(auth, name).then(r => closing()).catch(console.log).finally(() => setLoading(false))
    }
    const openning = () => setOpen(true)
    return <>
        <Fab onClick={() => openning()} color="primary"><Add /></Fab>
        <Dialog onClose={() => closing()} open={open}>
            <DialogTitle>Add Category</DialogTitle>
            <DialogContent>
                <TextField disabled={loading} onChange={e => setName(e.target.value)} value={name} margin="normal" required fullWidth id="name" label="Name" name="name" autoComplete="name" autoFocus />
            </DialogContent>
            <DialogActions>
                <Button disabled={loading} onClick={() => saving()}>SAVE</Button>
            </DialogActions>
        </Dialog>
    </>
}