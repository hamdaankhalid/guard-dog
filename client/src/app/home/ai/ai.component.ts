import { Component, OnInit } from '@angular/core';
import { RegisteredModel } from 'src/app/models/registered-model';
import { ModelRegistryService } from 'src/app/services/model-registry.service';

@Component({
  selector: 'app-ai',
  templateUrl: './ai.component.html',
  styleUrls: ['./ai.component.scss']
})
export class AiComponent implements OnInit {

  registeredModels: RegisteredModel[] = [];

  isUploading = false;
  isDeleting = false;

  private modelToBeUploaded: File | null = null;

  constructor(private modelRegistryService: ModelRegistryService) { }

  ngOnInit(): void {
    this.fetchModels();
  }

  private fetchModels() {
    this.modelRegistryService.getModels().subscribe((models: RegisteredModel[]) => this.registeredModels = models);
  }

  setModel(event: any) {
      this.modelToBeUploaded = event.target.files[0];
  }

  async registerModel() {
    if (!this.modelToBeUploaded) {
      return;
    }
    this.isUploading = true;
    await this.modelRegistryService.registerModel(this.modelToBeUploaded);
    this.isUploading = false;
    this.fetchModels();
  }

  deleteModel(id: number) {
    this.isDeleting = true;
    this.modelRegistryService.deleteModel(id).subscribe((_) => {
      this.fetchModels();
      this.isDeleting = false;
    });
  }
}
