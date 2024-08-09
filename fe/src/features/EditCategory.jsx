import { Edit } from "@mui/icons-material"
import { Button, Dialog, DialogActions, DialogContent, DialogTitle, Fab, TextField } from "@mui/material"
import { useState } from "react"
import { editCategory } from "../api/be"
import { useSelector } from "react-redux"

export default function EditCategory({ dbData, onClose, disabled }) {
    const auth = useSelector(s => s.position.auth)
    const [open, setOpen] = useState(false), [name, setName] = useState(dbData.name), [loading, setLoading] = useState(false)
    const closing = () => {
        setOpen(false)
        onClose()
    }
    const saving = () => {
        setLoading(true)
        editCategory(auth, name, dbData.id).then(r => closing()).catch(console.log).finally(() => setLoading(false))
    }
    const openning = () => setOpen(true)
    return <>
        <Fab disabled={disabled} onClick={() => openning()} color="secondary"><Edit /></Fab>
        <Dialog onClose={() => closing()} open={open}>
            <DialogTitle>Edit Category</DialogTitle>
            <DialogContent>
                <TextField disabled={loading} onChange={e => setName(e.target.value)} value={name} margin="normal" required fullWidth id="name" label="Name" name="name" autoComplete="name" autoFocus />
            </DialogContent>
            <DialogActions>
                <Button disabled={loading} onClick={() => saving()}>SAVE</Button>
            </DialogActions>
        </Dialog>
    </>
}