
import { Link } from "react-router-dom"
import { ModeToggle } from "@/components/mode-toggle"

export default function Navbar() {
  return (
    <nav className="w-full px-6 py-4 bg-white dark:bg-gray-900 shadow-sm flex items-center justify-between">
      {/* Left side */}
      <div className="flex items-center gap-6">
        <Link
          to="/"
          className="text-lg font-semibold text-gray-900 dark:text-gray-100"
        >
          MyFoodApp
        </Link>

        <Link
          to="/"
          className="text-gray-700 dark:text-gray-300 hover:underline"
        >
          Login
        </Link>

        <Link
          to="/register"
          className="text-gray-700 dark:text-gray-300 hover:underline"
        >
          Register
        </Link>

        <Link
          to="/order"
          className="text-gray-700 dark:text-gray-300 hover:underline"
        >
          Order Food
        </Link>
      </div>

      {/* Right side */}
      <ModeToggle />
    </nav>
  )
}

