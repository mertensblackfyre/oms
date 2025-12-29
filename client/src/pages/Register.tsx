import { Card, CardHeader, CardTitle, CardContent } from "@/components/ui/card"
import { Label } from "@/components/ui/label"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"

export default function Register() {
  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100 dark:bg-gray-900">
      <Card className="w-[380px] bg-white dark:bg-gray-800">
        <CardHeader>
          <CardTitle className="text-center text-gray-900 dark:text-gray-100">
            Create Account
          </CardTitle>
        </CardHeader>

        <CardContent>
          <form className="space-y-4">
            <div>
              <Label htmlFor="name" className="text-gray-900 dark:text-gray-100">
                Full Name
              </Label>
              <Input id="name" type="text" placeholder="John Doe" />
            </div>

            <div>
              <Label htmlFor="email" className="text-gray-900 dark:text-gray-100">
                Email
              </Label>
              <Input id="email" type="email" placeholder="you@example.com" />
            </div>

            <div>
              <Label htmlFor="password" className="text-gray-900 dark:text-gray-100">
                Password
              </Label>
              <Input id="password" type="password" placeholder="••••••••" />
            </div>

            <Button className="w-full" type="submit">
              Register
            </Button>
          </form>
        </CardContent>
      </Card>
    </div>
  )
}

