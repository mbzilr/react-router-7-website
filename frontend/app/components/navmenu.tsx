import {
  NavigationMenu,
  NavigationMenuList,
  NavigationMenuItem,
} from "~/components/ui/navigation-menu";
import { Link } from "react-router";

export default function NavMenu() {
  return (
    <div className="p-2 m-2">
      <NavigationMenu>
        <NavigationMenuList className="flex gap-4">
          {" "}
          {/* Add gap */}
          <NavigationMenuItem>
            <Link to="/" className="px-3 py-2">
              Home
            </Link>{" "}
            {/* Add padding */}
          </NavigationMenuItem>
          <NavigationMenuItem>
            <Link to="/about" className="px-3 py-2">
              About
            </Link>
          </NavigationMenuItem>
          <NavigationMenuItem>
            <Link to="/talents" className="px-3 py-2">
              Talents
            </Link>
          </NavigationMenuItem>
        </NavigationMenuList>
      </NavigationMenu>
    </div>
  );
}
