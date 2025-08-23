import { Button } from '../ui/button'
import { Trash2 } from 'lucide-react'
import { Dialog, DialogClose, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from '../ui/dialog'

export default function DeleteBook() {
  return (
    <div>
      <div className="flex flex-col space-y-3">
        <div className="flex items-center justify-between">
          <p className="text-2xl font-bold text-primary">Danger Zone</p>
          <Dialog>
            <DialogTrigger asChild>
              <Button variant="destructive" className="flex items-center">
                <Trash2 className="h-4 w-4" />
                <span>Delete Book</span>
              </Button>
            </DialogTrigger>
            <DialogContent className="sm:max-w-md">
              <DialogHeader>
                <DialogTitle>Confirm Deletion</DialogTitle>
                <DialogDescription>
                  Are you sure you want to delete this book? This action cannot be undone.
                </DialogDescription>
              </DialogHeader>

              <DialogFooter className="sm:justify-end">
                <DialogClose asChild>
                  <Button type="button" variant="secondary">
                    Cancel
                  </Button>
                </DialogClose>
                <Button type="button" variant="destructive">
                  Confirm Delete
                </Button>
              </DialogFooter>
            </DialogContent>
          </Dialog>
        </div>
      </div>
    </div>
  )
}
