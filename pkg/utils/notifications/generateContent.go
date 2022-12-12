package notifications

import "fmt"

func GenerateContent(title, name string, invoiceID int) string {
	switch title {
	// case "Payment is Overdue":
	// 	return fmt.Sprintf("Your bill from %s invoice #%d is %d days overdue. Please make a payment immediately",)
	// case "Payment is Due Soon":
	// 	return fmt.Sprintf("Your bill from %s invoice #%d is %d days overdue. Please make a payment immediately",)
	case "Payment Success":
		return fmt.Sprintf("Your payment to %s for invoice #%d was successful!", name, invoiceID)
	case "Payment Failed":
		return fmt.Sprintf("Your payment to %s for invoice #%d was failed for some reason. Please try again!", name, invoiceID)
	case "Payment Pending":
		return fmt.Sprintf("Your payment to %s for invoice #%d is waiting to be checked", name, invoiceID)
	case "Payment Done":
		return fmt.Sprintf("%s has paid for invoice #%d. Check payment now!", name, invoiceID)
	case "New Bill Issued":
		return fmt.Sprintf("%s issued you a bill with invoice #%d check your invoice now!", name, invoiceID)
	}

	return ""
}
