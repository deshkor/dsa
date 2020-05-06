import java.util.Arrays;
import java.util.Scanner;
import java.util.ArrayList;
import java.util.List;
import java.io.File;
import java.io.FileNotFoundException;

class BinarySearch {
	private BinarySearch() {}

	public static int binSearch(int []a, int key) {
		int lower = 0;
		int higher = a.length - 1;

		while (lower <= higher) {
			int middle = lower + (higher - lower) / 2;

			if (key < a[middle]) higher = middle - 1;
			else if (key > a[middle]) lower = middle + 1;
			else return middle;
		}
		return -1;
	}

	public static void main(String[] args) {
		if (args.length != 1) {
			System.out.println("Needs a file read.");
			System.exit(1);
		}

		try {
			Scanner scanner = new Scanner(new File(args[0]));
			List<Integer> list = new ArrayList<Integer>();

			while (scanner.hasNextInt()) {
				list.add(scanner.nextInt());
			}

			int[] listArray = list.stream().mapToInt(i->i).toArray();
			Arrays.sort(listArray);

			for (int i = 0; i < listArray.length; i ++) {
				System.out.printf("[%d] %d\n", i, listArray[i]);
			}

			int idx;
			int key;
			Scanner in = new Scanner(System.in);

			System.out.printf("Enter a number or and array of numbers [ctrl+c to exit]: ");
			while (in.hasNextInt()) {
				key = in.nextInt();
				idx = BinarySearch.binSearch(listArray, key);
				if (idx != -1) {
					System.out.printf("Found %d at index %d.\n", key, idx);
				} else {
					System.out.printf("Couldn't find number %d in array.\n", key);
				}

				System.out.printf("Enter a number or and array of numbers [ctrl+c to exit]: ");
			}
		} catch (FileNotFoundException e) {
			System.out.println("FileNotFoundException: " + e.getMessage());
			System.exit(1);
		}
	}
}
