import { Card, CardHeader, CardTitle, CardContent } from "@/components/ui/card"
import { Label } from "@/components/ui/label"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"

export default function Login() {
  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100 dark:bg-gray-900">
      <Card className="w-[380px] bg-white dark:bg-gray-800">
        <CardHeader>
          <CardTitle className="text-center text-gray-900 dark:text-gray-100">
            Login
          </CardTitle>
        </CardHeader>

        <CardContent>
          <form className="space-y-4">
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
              Sign In
            </Button>
          </form>
        </CardContent>
      </Card>
    </div>
  )
}

