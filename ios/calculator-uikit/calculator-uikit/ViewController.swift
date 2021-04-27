//
//  ViewController.swift
//  calculator-uikit
//
//  Created by Anton Kapralov on 26.04.2021.
//

import UIKit

class ViewController: UIViewController {

  @IBOutlet weak var leftTextField: UITextField!
  @IBOutlet weak var rightTextField: UITextField!
  @IBOutlet weak var sumTextField: UITextField!
  
  override func viewDidLoad() {
    super.viewDidLoad()
    // Do any additional setup after loading the view.
  }
  
  @IBAction func onCalculate(_ sender: Any) {
    let left: Int = Int(leftTextField.text ?? "") ?? 0
    let right: Int = Int(rightTextField.text ?? "") ?? 0

    sumTextField.text = String(left + right)
  }
  
}

